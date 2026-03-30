# Roadmap — whatsapp-go-sdk

**Referencia:** [MASTER_BLUEPRINT.md](../MASTER_BLUEPRINT.md)
**Ultima atualizacao:** 2026-03-28

---

## Estado Atual: v1.0.1 (Estavel)

Todos os 9 modulos originais completos com 100% de cobertura de testes.

---

## TASK 0 (OBRIGATORIA): State Checkpoint

**Antes de iniciar QUALQUER milestone ou task, o agente DEVE:**

1. Ler `.agents/STATE_CHECKPOINT.json` deste repositorio
2. Comparar `last_checkpoint.commit_hash` com `git log --oneline -1`
3. Se hash diferente → executar `git log --oneline <old>..HEAD` e `git diff --stat <old>..HEAD`
4. Classificar mudancas como COMPATIVEL / RELEVANTE / CONFLITANTE
5. Se CONFLITANTE → reportar ao usuario antes de prosseguir
6. Se dirty working tree (`git status --short` nao vazio) → reportar ao usuario
7. Ao concluir cada task → atualizar `STATE_CHECKPOINT.json` com novo commit hash

**Detalhes completos:** Ver [MASTER_BLUEPRINT.md — Task 0](../MASTER_BLUEPRINT.md#task-0-obrigatoria-state-checkpoint--verificacao-de-estado-pre-execucao)

**Arquivo de controle:** `.agents/STATE_CHECKPOINT.json`

### TASK 0.5 (OBRIGATORIA): Refinamento Pre-Milestone

**Apos o checkpoint e ANTES de executar qualquer task, o agente DEVE:**

1. Ler os arquivos deste repo que serao impactados pelo milestone
2. Listar o que JA EXISTE vs o que PRECISA SER CRIADO
3. Decompor cada task em sub-tasks atomicas (1-3 arquivos por sub-task)
4. Sequenciar: domain → ports → services → adapters → usecases → testes
5. Apresentar o plano ao usuario antes de comecar a codar
6. Salvar o plano em `.agents/REFINEMENT_M<N>.md`

**Cada sub-task = 1 commit. TDD obrigatorio (teste antes da implementacao).**

**Detalhes completos:** Ver [MASTER_BLUEPRINT.md — Secao 0.5](../MASTER_BLUEPRINT.md#05-refinamento-obrigatorio-pre-milestone-atomizacao-de-sub-tasks)

---

## Tarefas Pendentes (por Milestone)

### MILESTONE 3: Extensoes para SaaS

| # | Task | Descricao | Status |
|---|------|-----------|--------|
| 3.1 | Mark as Read | `Messages().MarkAsRead(ctx, messageID)` — POST /{phoneID}/messages com `status: "read"` | PENDENTE |
| 3.2 | Send Video/Document/Audio/Location | Metodos diretos `SendVideo()`, `SendDocument()`, `SendAudio()`, `SendLocation()` no MessageService | PENDENTE |
| 3.3 | Interactive Lists | `SendInteractiveList()` no MessageService — lista de opcoes para o usuario | PENDENTE |
| 3.4 | Reaction Messages | `SendReaction()` no MessageService — reagir a mensagem com emoji | PENDENTE |

### Regras de Implementacao

- TDD obrigatorio (testes ANTES da implementacao)
- Manter zero dependencias externas
- Seguir padrao hexagonal existente (domain -> ports -> services -> adapters -> usecases)
- Cada feature deve ter builder no usecases/ e metodo no service
- Atualizar `client_test.go` para cobrir novos metodos

### Notas Tecnicas

**3.1 Mark as Read:** Endpoint Meta: `POST /{phone_number_id}/messages` com body:
```json
{
    "messaging_product": "whatsapp",
    "status": "read",
    "message_id": "wamid.xxx"
}
```

**3.2 Media Messages:** Ja existem builders parciais nos usecases. Falta expor como metodos diretos no `MessageService` interface e implementacao.

**3.3 Interactive Lists:** Similar ao `SendInteractiveButton`, mas com type `"list"` e sections com rows.

**3.4 Reactions:** Endpoint Meta: POST messages com type `"reaction"` e body `{ message_id, emoji }`.
