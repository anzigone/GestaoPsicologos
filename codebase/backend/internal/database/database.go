package database

import (
	"crypto/rand"
	"database/sql"
	"fmt"
	"os"
	"path/filepath"

	"github.com/anzigone/GestaoPsicologos/backend/internal/config"
	_ "github.com/go-sql-driver/mysql"
	_ "modernc.org/sqlite"
)

func Connect(cfg *config.Config) (*sql.DB, error) {
	switch cfg.DBDriver {
	case "sqlite":
		dir := filepath.Dir(cfg.DatabasePath)
		if err := os.MkdirAll(dir, 0755); err != nil {
			return nil, fmt.Errorf("falha ao criar diretório do banco: %w", err)
		}
		db, err := sql.Open("sqlite", cfg.DatabasePath)
		if err != nil {
			return nil, fmt.Errorf("falha ao abrir SQLite: %w", err)
		}
		if err := db.Ping(); err != nil {
			return nil, fmt.Errorf("falha ao conectar ao SQLite: %w", err)
		}
		return db, nil

	case "mysql":
		dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s?parseTime=true&charset=utf8mb4",
			cfg.DBUser, cfg.DBPassword, cfg.DBHost, cfg.DBName)
		db, err := sql.Open("mysql", dsn)
		if err != nil {
			return nil, fmt.Errorf("falha ao conectar ao MySQL: %w", err)
		}
		if err := db.Ping(); err != nil {
			return nil, fmt.Errorf("falha ao conectar ao MySQL: %w", err)
		}
		return db, nil

	default:
		return nil, fmt.Errorf("driver de banco não suportado: %s", cfg.DBDriver)
	}
}

// Seed inserts the admin master user if the users table is empty.
// passwordHash is a pre-computed SHA256+salt hash from the auth package.
func Seed(db *sql.DB, passwordHash string) error {
	var count int
	if err := db.QueryRow("SELECT COUNT(*) FROM users").Scan(&count); err != nil {
		return fmt.Errorf("falha ao verificar tabela users: %w", err)
	}
	if count > 0 {
		return nil
	}
	b := make([]byte, 16)
	_, _ = rand.Read(b)
	b[6] = (b[6] & 0x0f) | 0x40
	b[8] = (b[8] & 0x3f) | 0x80
	id := fmt.Sprintf("%x-%x-%x-%x-%x", b[0:4], b[4:6], b[6:8], b[8:10], b[10:])

	_, err := db.Exec(
		`INSERT INTO users (id, email, password_hash, role, name) VALUES (?, ?, ?, ?, ?)`,
		id, "admin@admin.com.br", passwordHash, "admin", "Administrador Master",
	)
	return err
}

// Migrate executa a criação das tabelas base.
// ON UPDATE CURRENT_TIMESTAMP é omitido por compatibilidade SQLite/MySQL;
// o campo updated_at é atualizado explicitamente na camada de repositório.
func Migrate(db *sql.DB, driver string) error {
	statements := []string{
		`CREATE TABLE IF NOT EXISTS users (
			id           VARCHAR(36)    PRIMARY KEY,
			email        VARCHAR(255)   UNIQUE NOT NULL,
			password_hash VARCHAR(255)  NOT NULL,
			role         VARCHAR(20)    NOT NULL,
			name         VARCHAR(255)   NOT NULL,
			crp          VARCHAR(50),
			specialty    VARCHAR(100),
			phone        VARCHAR(20),
			base_fee     DECIMAL(10,2)  DEFAULT 0.00,
			package_sessions INT,
			package_fee  DECIMAL(10,2),
			created_at   DATETIME       DEFAULT CURRENT_TIMESTAMP,
			updated_at   DATETIME       DEFAULT CURRENT_TIMESTAMP
		)`,
		`CREATE TABLE IF NOT EXISTS oauth_tokens (
			user_id       VARCHAR(36) NOT NULL,
			provider      VARCHAR(20) NOT NULL,
			access_token  TEXT        NOT NULL,
			refresh_token TEXT,
			expiry        DATETIME    NOT NULL,
			updated_at    DATETIME    DEFAULT CURRENT_TIMESTAMP,
			PRIMARY KEY (user_id, provider),
			FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE
		)`,
		`CREATE TABLE IF NOT EXISTS active_sessions (
			id         VARCHAR(36) PRIMARY KEY,
			user_id    VARCHAR(36) NOT NULL,
			token      TEXT        NOT NULL,
			expires_at DATETIME    NOT NULL,
			created_at DATETIME    DEFAULT CURRENT_TIMESTAMP,
			FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE
		)`,
		`CREATE TABLE IF NOT EXISTS patients (
			id               VARCHAR(36)   PRIMARY KEY,
			psychologist_id  VARCHAR(36)   NOT NULL,
			name             VARCHAR(255)  NOT NULL,
			phone            VARCHAR(20),
			birthdate        VARCHAR(10),
			age              INT,
			profession       VARCHAR(100),
			company          VARCHAR(100),
			city             VARCHAR(100),
			state            VARCHAR(2),
			marital_status   VARCHAR(30),
			consultation_fee DECIMAL(10,2) DEFAULT 0.00,
			active           INTEGER       NOT NULL DEFAULT 1,
			created_at       DATETIME      DEFAULT CURRENT_TIMESTAMP,
			updated_at       DATETIME      DEFAULT CURRENT_TIMESTAMP,
			FOREIGN KEY (psychologist_id) REFERENCES users(id) ON DELETE CASCADE
		)`,
		`CREATE TABLE IF NOT EXISTS first_analysis (
			patient_id              VARCHAR(36) PRIMARY KEY,
			main_complaint          TEXT,
			symptom_diagnosis       TEXT,
			developmental_influence TEXT,
			situational_issues      TEXT,
			biological_factors      TEXT,
			strengths_resources     TEXT,
			addictions              TEXT,
			stimuli                 TEXT,
			thoughts                TEXT,
			behaviors               TEXT,
			affects                 TEXT,
			physiological           TEXT,
			treatment_goals         TEXT,
			treatment_plan          TEXT,
			updated_at              DATETIME    DEFAULT CURRENT_TIMESTAMP,
			FOREIGN KEY (patient_id) REFERENCES patients(id) ON DELETE CASCADE
		)`,
		`CREATE TABLE IF NOT EXISTS sessions (
			id               VARCHAR(36)  PRIMARY KEY,
			patient_id       VARCHAR(36)  NOT NULL,
			session_date     DATETIME     NOT NULL,
			notes            TEXT,
			status           VARCHAR(20)  NOT NULL DEFAULT 'pendente',
			meet_link        VARCHAR(500),
			outlook_event_id VARCHAR(255),
			created_at       DATETIME     DEFAULT CURRENT_TIMESTAMP,
			updated_at       DATETIME     DEFAULT CURRENT_TIMESTAMP,
			FOREIGN KEY (patient_id) REFERENCES patients(id) ON DELETE CASCADE
		)`,
	}

	for _, stmt := range statements {
		if _, err := db.Exec(stmt); err != nil {
			return fmt.Errorf("falha na migração: %w", err)
		}
	}

	// Additive migrations: add columns to existing tables without version tracking.
	if err := migrateAddColumn(db, driver, "patients", "active", "INTEGER NOT NULL DEFAULT 1"); err != nil {
		return fmt.Errorf("falha ao adicionar coluna active: %w", err)
	}
	return nil
}

func migrateAddColumn(db *sql.DB, driver, table, column, definition string) error {
	var exists int
	switch driver {
	case "mysql":
		db.QueryRow(
			`SELECT COUNT(*) FROM INFORMATION_SCHEMA.COLUMNS WHERE TABLE_NAME=? AND COLUMN_NAME=? AND TABLE_SCHEMA=DATABASE()`,
			table, column,
		).Scan(&exists)
	default: // sqlite
		db.QueryRow(
			fmt.Sprintf(`SELECT COUNT(*) FROM pragma_table_info('%s') WHERE name=?`, table), column,
		).Scan(&exists)
	}
	if exists > 0 {
		return nil
	}
	_, err := db.Exec(fmt.Sprintf(`ALTER TABLE %s ADD COLUMN %s %s`, table, column, definition))
	return err
}
