# Meta 3: Concorrência Massiva

Passos para adicionar comportamento concorrente ao jogo usando goroutines e canais.

## Passo 1 – Criar goroutines de bot
1. Para cada bot, inicie uma goroutine separada:
   ```go
   go runBot(bot, actionChan, stopChan)
   ```
2. O bot deve ler um ticker e enviar comandos de movimento/ataque.

## Passo 2 – Criar canais de comunicação
1. Defina um tipo de mensagem:
   ```go
   type Action struct {
       entityID string
       actionType string
       dx, dy float64
   }
   ```
2. Crie o canal central:
   ```go
   actionChan := make(chan Action)
   ```
3. Todas as decisões de entidades entram nesse canal.

## Passo 3 – Tickers assíncronos
1. Em cada bot ou entidade, use um `time.Ticker`:
   ```go
   ticker := time.NewTicker(200 * time.Millisecond)
   defer ticker.Stop()

   for {
       select {
       case <-ticker.C:
           // tomar decisão
       case <-stopChan:
           return
       }
   }
   ```
2. Isso cria cadências de decisão independentes.

## Passo 4 – Separar intenção de execução
1. Use o canal para enviar intenções, mas não alterar estado diretamente.
2. Um árbitro central processa as ações e atualiza a posição/vida.

## Passo 5 – Sincronizar a atualização global
1. Adicione um loop de processamento central:
   ```go
   go func() {
       for action := range actionChan {
           processAction(action)
       }
   }()
   ```
2. Nunca modifique estado compartilhado fora desse loop.

## Resultado esperado
- Cada bot age em sua própria goroutine.
- A lógica de movimento/ataque é enviada por canal.
- O jogo passa a usar CSP em vez de estado global direto.

## Próximo passo
Com a concorrência do bot pronta, implemente combate e projéteis funcionais.
