# Sprint 10: Hardening e Prontidão para Produção (sprint10.md)

## 1. Escopo e Objetivos
O objetivo desta sprint é reforçar a segurança do sistema antes da implantação na nuvem e assegurar a responsividade em celulares. A autenticação será migrada para cookies protegidos contra roubo de sessão, e habilitaremos o suporte para leitura de credenciais no AWS Secrets Manager.

**Objetivos principais**:
- Migrar o armazenamento e envio de tokens JWT para cookies `HttpOnly`, `Secure` e `SameSite`.
- Atualizar middlewares no backend e frontend para suportar a autenticação baseada em cookies.
- Configurar o módulo de leitura de parâmetros seguros no AWS Secrets Manager/Parameter Store para produção.
- Desenvolver um menu móvel funcional (Drawer/Sheet) para telas compactas.

---

## 2. Entregáveis da Sprint
Ao final desta sprint, os seguintes entregáveis deverão estar prontos:
- Sessão do usuário autenticada por meio de cookies HTTP-only seguros e transparentes ao código cliente.
- Mecanismo de leitura de chaves de API e banco de dados via SDK da AWS integrado ao boot da aplicação Go.
- Componente de Menu Mobile (hambúrguer) responsivo ativo no frontend para acesso à Sidebar de navegação.

---

## 3. Checkpoints de Validação Humana
Para validar a conclusão desta Sprint, o humano realizará os seguintes testes:

1. **Subir os serviços**:
   - Executar na raiz do projeto:
     ```bash
     docker compose up --build -d
     ```
2. **Inspeção de Segurança de Cookies**:
   - Efetuar login e abrir o painel de desenvolvedor (F12) na aba de Cookies de `http://localhost:3000`.
   - Esperado: O cookie contendo o token JWT está com as opções **HttpOnly** e **Secure** marcadas (em ambiente local HTTP, a flag `Secure` pode estar desabilitada temporariamente para testes se o navegador exigir, mas deve estar habilitada em produção).
3. **Impedimento de Acesso Client-side**:
   - Abrir o console JavaScript do navegador e digitar:
     ```javascript
     document.cookie
     ```
   - Esperado: O token de autenticação não é retornado na string de saída (invisibilidade garantindo segurança contra ataques XSS).
4. **Validação Mobile**:
   - Alternar a visualização para modo celular no navegador (largura menor que 768px).
   - Esperado: A barra de navegação padrão desaparece, dando lugar ao botão hambúrguer. Clicar no botão abre a gaveta de links laterais com usabilidade fluida.
5. **Encerramento da Sprint**:
   - Realizar commit com a mensagem `feat: sprint 10 hardening concluido`, criar a tag Git `v0.10.0` e fazer push para a branch `main`.
 - fazer o push https://github.com/anzigone/GestaoPsicologos.git