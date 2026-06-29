# Sprint 2: Webservice BFF Mockado e Documentado (sprint2.md)

## 1. Escopo e Objetivos
O objetivo desta sprint é implementar a totalidade dos endpoints do webservice backend (BFF) em Golang de forma mockada (respostas simuladas com dados estáticos realistas) e gerar a documentação interativa através do Swagger/OpenAPI. Isso permitirá o desenvolvimento independente do Frontend nas próximas sprints.

**Objetivos principais**:
- Configurar o gerador e leitor de documentação Swagger (`swaggo/swag`) na API.
- Criar a árvore completa de rotas HTTP REST.
- Implementar as respostas fictícias (JSON mockado) para todos os fluxos de negócios baseados nas especificações.
- Garantir o tratamento correto do CORS para a comunicação futura com o frontend Next.js.

---

## 2. Entregáveis da Sprint
Ao final desta sprint, os seguintes entregáveis deverão estar prontos:
- Documentação interativa da API exposta publicamente em `http://localhost:8080/swagger/index.html`.
- Roteador HTTP completo contendo todas as rotas de autenticação, administração master, pacientes, primeira análise, sessões e integrações, devolvendo as assinaturas e status HTTP corretos.
- Implementação mockada no backend (sem persistência real em banco nesta etapa) retornando respostas padrão em Português do Brasil (PT-BR).

---

## 3. Checkpoints de Validação Humana
Para validar a conclusão desta Sprint, o humano realizará os seguintes testes:

1. **Subir os serviços**:
   - Executar na raiz do projeto:
     ```bash
     docker compose up --build -d
     ```
2. **Acessar Documentação Swagger**:
   - Navegar até `http://localhost:8080/swagger/index.html`.
   - Esperado: Tela do Swagger UI listando detalhadamente todas as rotas sob as tags `Autenticação`, `Administração`, `Pacientes`, `Primeira Análise`, `Sessões` e `Integrações`.
3. **Validar Respostas REST**:
   - Executar uma chamada de teste no Swagger no método `POST /api/auth/login`.
     - Esperado: Retorno `200 OK` contendo um token JWT fictício e a estrutura de dados do usuário psicólogo.
   - Executar uma chamada de teste no Swagger no método `GET /api/patients`.
     - Esperado: Retorno `200 OK` contendo uma lista JSON simulada de pacientes cadastrados.
4. **Encerramento da Sprint**:
   - Realizar commit com a mensagem `feat: sprint 2 mock webservice concluido`, criar a tag Git `v0.2.0` e fazer push para a branch `main`.
