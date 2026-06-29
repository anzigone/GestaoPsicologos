# Sprint 5: Administração do Psicólogo (sprint5.md)

## 1. Escopo e Objetivos
O objetivo desta sprint é implementar a persistência de dados cadastrais, tarifas e controle de acesso dos psicólogos. Toda a lógica de leitura e gravação dos dados do profissional no banco de dados será implementada, junto com a troca de senha.

**Objetivos principais**:
- Implementar o repositório e endpoints do perfil do psicólogo no backend.
- Assegurar que as alterações de perfil não permitam violação de privilégios de outros psicólogos (isolamento por ID).
- Integrar a interface do formulário de administração com a API REST.
- Viabilizar a alteração de senha segura com verificação de senha antiga no banco.

---

## 2. Entregáveis da Sprint
Ao final desta sprint, os seguintes entregáveis deverão estar prontos:
- Endpoints funcionais no backend:
  - `GET /api/psychologist`: Retorna os dados do psicólogo logado.
  - `PUT /api/psychologist`: Atualiza os dados profissionais e valores de consulta.
  - `POST /api/auth/change-password`: Valida a senha atual e define a nova senha criptografada.
- Integração da coluna esquerda da página `/admin/settings` no frontend para edição cadastral e troca de senha do psicólogo logado.

---

## 3. Checkpoints de Validação Humana
Para validar a conclusão desta Sprint, o humano realizará os seguintes testes:

1. **Subir os serviços**:
   - Executar na raiz do projeto:
     ```bash
     docker compose up --build -d
     ```
2. **Atualização Cadastral**:
   - Logar com uma conta de psicólogo criada anteriormente.
   - Navegar até a página `/admin/settings`.
   - Modificar os campos: Especialidade para "Terapia Cognitivo Comportamental", Valor Padrão da Consulta para R$ 180,00 e telefone celular. Clicar em "Salvar".
   - Esperado: Toast ou mensagem de sucesso em PT-BR.
3. **Persistência Física**:
   - Recarregar a página `/admin/settings` (F5).
   - Esperado: Os campos recarregam preenchidos com os novos valores diretamente do banco SQLite.
4. **Troca de Senha**:
   - No formulário inferior de "Troca de Senha", digitar a senha atual correta e a nova senha. Clicar em "Alterar Senha".
   - Esperado: Mensagem de confirmação.
   - Realizar logout e tentar logar com a nova senha definida.
   - Esperado: Login bem-sucedido.
5. **Encerramento da Sprint**:
   - Realizar commit com a mensagem `feat: sprint 5 administracao do psicologo concluido`, criar a tag Git `v0.5.0` e fazer push para a branch `main`.
