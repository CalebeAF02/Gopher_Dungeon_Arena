# Meta 5: Sincronização e Estado Seguro

Passos para proteger o estado do jogo e garantir atualizações consistentes.

## Passo 1 – Proteger `HP` com mutex
1. Adicione um campo `sync.Mutex` em entidades cujo estado pode mudar:
   ```go
   type BaseEntity struct {
       mu sync.Mutex
       health int
       // outros campos
   }
   ```
2. Encapsule leituras e escritas:
   ```go
   func (e *BaseEntity) Damage(amount int) {
       e.mu.Lock()
       defer e.mu.Unlock()
       e.health -= amount
   }
   ```

## Passo 2 – Criar um árbitro global
1. O árbitro recebe ações e decide se são válidas.
2. Ele processa apenas uma ação por vez e atualiza o estado oficial.

## Passo 3 – Usar canais como única via de alteração
1. Nunca modifique `health`, `position` ou contadores diretamente fora do árbitro.
2. Todas as intenções de movimento e ataque devem passar por `actionChan`.

## Passo 4 – Adicionar logs de sistema
1. Imprima mensagens para eventos importantes:
   - nova goroutine de bot criada
   - bot morto
   - jogador atingido
   ```go
   log.Printf("Bot %s entrou no jogo", bot.ID)
   ```
2. Logs no terminal ajudam a entender a concorrência.

## Passo 5 – Verificar condições de corrida
1. Execute o programa com `go run -race ./src`.
2. Corrija qualquer condição de corrida relatada.

## Resultado esperado
- Estado crítico é protegido.
- A lógica principal roda em sequência no árbitro.
- O jogo não tem condições de corrida visíveis.

## Próximo passo
Com o estado seguro, adicione a dinâmica da dungeon para criar desafios e objetivos.
