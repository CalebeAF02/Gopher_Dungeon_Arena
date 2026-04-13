# Meta 4: Mecânicas de Combate e Projéteis

Passos para criar disparos, colisões e feedback de dano no jogo.

## Passo 1 – Criar projéteis em goroutines
1. Defina a struct `Projectile` com posição, direção e velocidade.
2. Ao atirar, inicie uma goroutine que move o projétil:
   ```go
   func runProjectile(p *Projectile, world *World, done chan<- struct{}) {
       ticker := time.NewTicker(16 * time.Millisecond)
       defer ticker.Stop()
       for range ticker.C {
           p.x += p.dx
           p.y += p.dy
           if collision := world.CheckCollision(p); collision {
               done <- struct{}{}
               return
           }
       }
   }
   ```

## Passo 2 – Detectar colisões com AABB
1. Use colisão de caixa alinhada (AABB) para entidades:
   ```go
   func aabbOverlap(x1, y1, w1, h1, x2, y2, w2, h2 float64) bool {
       return x1 < x2+w2 && x1+w1 > x2 && y1 < y2+h2 && y1+h1 > y2
   }
   ```
2. Compare o projétil com cada entidade.

## Passo 3 – Aplicar dano e remover projéteis
1. Quando uma colisão ocorre, envie um evento de dano ao árbitro.
2. Atualize `health` da entidade atingida e marque o projétil para remoção.

## Passo 4 – Feedback visual de dano
1. Ao sofrer um hit, altere temporariamente a cor da entidade.
2. Use um contador curto (`damageFlashFrames`) para voltar à cor normal.

## Passo 5 – Testar a lógica de combate
1. Faça dois objetos simulados e um projétil.
2. Verifique se a colisão reduz `health` e remove o projétil.
3. Confirme se a tela exibe o flash de dano.

## Resultado esperado
- Projéteis se movem em linha.
- Colisões com bots/jogadores são detectadas.
- Dano é aplicado de forma visível.

## Próximo passo
Agora que o combate funciona, proteja o estado compartilhado e implemente a sincronização correta.
