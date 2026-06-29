# Sprint 9: Polimento e Ajustes Gerais (sprint9.md)

## 1. Escopo e Objetivos
O objetivo desta sprint é elevar o padrão de usabilidade e acabamento do sistema, realizando um refactoring focado na experiência do usuário (UX/UI), validações de inputs, formatação de dados e correções de timezone em toda a aplicação.

**Objetivos principais**:
- Adicionar estados visuais de carregamento (spinners, skeletons) em ações assíncronas.
- Implementar validação robusta de formulários (Zod e React Hook Form) no frontend e tratamento correspondente no backend.
- Padronizar a exibição de moedas (R$) e datas no padrão brasileiro (PT-BR).
- Resolver inconsistências de fusos horários no salvamento e leitura de sessões/agendamentos.

---

## 2. Entregáveis da Sprint
Ao final desta sprint, os seguintes entregáveis deverão estar prontos:
- Validação e tratamento de dados implementados em todos os formulários.
- Componentes visuais do Shadcn/UI com feedback de carregamento em execução.
- Normalização de Timezone (UTC) configurada no backend e formatadores de data/moeda unificados no frontend.

---

## 3. Checkpoints de Validação Humana
Para validar a conclusão desta Sprint, o humano realizará os seguintes testes:

1. **Subir os serviços**:
   - Executar na raiz do projeto:
     ```bash
     docker compose up --build -d
     ```
2. **Validação de Formulários**:
   - Tentar salvar um novo paciente sem preencher o telefone ou informando um e-mail em formato incorreto.
   - Esperado: Ação impedida e mensagens de erro explícitas em vermelho sob os inputs correspondentes em PT-BR.
3. **Feedback de Interface**:
   - Executar operações de salvamento (Cadastro do Psicólogo, Anamnese, etc.).
   - Esperado: Botões desabilitados exibindo estado de "Salvando..." ou um spinner enquanto a chamada à API é processada.
4. **Alinhamento de Fusos**:
   - Criar uma sessão em um horário específico, recarregar a tela e conferir se o horário exibido na timeline bate exatamente com o cadastrado.
5. **Encerramento da Sprint**:
   - Realizar commit com a mensagem `feat: sprint 9 polimento e ajustes concluido`, criar a tag Git `v0.9.0` e fazer push para a branch `main`.
