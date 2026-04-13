# Meta 7: Interface e Input de Utilizador

Passos para adicionar controles e interface visual ao jogo.

## Passo 1 – Configurar multi-input
1. Use `ebiten.IsKeyPressed` para mapear teclas.
2. Exemplo de controles locais:
   - Jogador 1: `W`, `A`, `S`, `D`
   - Jogador 2: `Seta para cima`, `Seta para baixo`, `Seta para esquerda`, `Seta para direita`
3. No `Update`, leia as teclas e envie ações para o árbitro.

## Passo 2 – Criar um HUD básico
1. Desenhe barras de vida sobre cada jogador:
   ```go
   func drawHealthBar(screen *ebiten.Image, x, y float64, health int) {
       // desenhar fundo e preenchimento
   }
   ```
2. Use cores diferentes para representar vida restante.
3. Mostre também o nome do time ou número do jogador.

## Passo 3 – Implementar telas de estado
1. Crie estados de jogo: `Playing`, `Victory`, `GameOver`.
2. No `Draw`, desenhe mensagens grandes quando o jogo terminar:
   - "Game Over"
   - "Vitória do Time A"
   - "Vitória do Time B"
3. Adicione contagem de membros restantes:
   ```go
   fmt.Sprintf("Vivos: %d", world.AliveCount())
   ```

## Passo 4 – Feedback visual imediato
1. Mostre `+` ou `-` de dano ao receber hits.
2. Desenhe o número de bots restantes ou o tempo decorrido.

## Passo 5 – Testar a interface
1. Verifique se os controles respondem sem travamentos.
2. Confirme se o HUD é legível.
3. Confirme se as telas de vitória/derrota aparecem corretamente.

## Resultado esperado
- Jogadores controlam entidades localmente.
- O jogo indica vida e estado atual.
- Mensagens de fim de partida são claras.

## Próximo passo
Finalize a solução acadêmica com cancelamento limpo e tratamento de erros.
