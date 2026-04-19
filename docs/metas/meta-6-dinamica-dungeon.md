# Meta 6: Dinâmica da Dungeon

Passos para construir a jogabilidade da dungeon e os desafios do mapa.

## Passo 1 – Spawn progressivo de bots
1. Use um `time.Ticker` ou contador de tempo para aumentar bots ativos.
2. A cada intervalo, crie novos bots e inicie sua goroutine:
   ```go
   func spawnBotsPeriodically(world *World, interval time.Duration) {
       ticker := time.NewTicker(interval)
       defer ticker.Stop()
       for range ticker.C {
           world.AddBot()
       }
   }
   ```
3. Controle um limite máximo para não sobrecarregar o jogo.

## Passo 2 – Definir a zona de extração
1. Escolha uma área no final do mapa.
2. Verifique se um jogador vivo entrou na zona:
   ```go
   if insideExtractionZone(player.x, player.y) && player.Health() > 0 {
       world.EndGame(team)
   }
   ```
3. Marque a vitória ao menos de um membro chegar vivo.

## Passo 3 – Obstáculos estáticos
1. Defina paredes ou barreiras como retângulos fixos.
2. No movimento, impeça passagem quando houver interseção.
3. Para projéteis, faça o mesmo teste de colisão com as paredes.

## Passo 4 – Mapear o sistema de coordenadas
1. Crie uma matriz simples ou lista de blocos do mapa.
2. Cada bloco pode ser livre ou bloqueado.
3. Use essa estrutura para guiar movimento e renderização.

## Passo 5 – Testar a jogabilidade
1. Rode o jogo e observe a quantidade de bots aumentando.
2. Tente chegar à extração.
3. Verifique se paredes bloqueiam movimento e tiros.

## Resultado esperado
- Dificuldade cresce ao longo do tempo.
- Há um objetivo claro: chegar à extração.
- Obstáculos adicionam estratégia ao movimento.

## Próximo passo
Depois desta dinâmica, implemente os controles do jogador e a HUD visual.
