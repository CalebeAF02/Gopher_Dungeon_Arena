# Meta 2: Entidades e Tipagem

Passos para criar a arquitetura básica de entidades no jogo usando tipos Go.

## Passo 1 – Definir uma interface `Entity`
1. Crie uma interface com métodos comuns:
   ```go
   type Entity interface {
       Position() (x, y float64)
       Health() int
       Team() string
       Update()
       Draw(screen *ebiten.Image)
   }
   ```
2. Esse contrato garante que todas as entidades possam ser tratadas de forma uniforme.

## Passo 2 – Criar structs principais
1. Defina structs para cada tipo de entidade:
   - `Player`
   - `Bot`
   - `Projectile`
2. Use composição para campos comuns:
   ```go
   type BaseEntity struct {
       x, y float64
       health int
       team string
   }

   type Player struct {
       BaseEntity
       speed float64
   }

   type Bot struct {
       BaseEntity
       aiState string
   }

   type Projectile struct {
       x, y float64
       dx, dy float64
       owner string
   }
   ```

## Passo 3 – Implementar os métodos da interface
1. Adicione métodos para `Position`, `Health` e `Team`:
   ```go
   func (e *BaseEntity) Position() (float64, float64) {
       return e.x, e.y
   }

   func (e *BaseEntity) Health() int {
       return e.health
   }

   func (e *BaseEntity) Team() string {
       return e.team
   }
   ```
2. Cada entidade deve ter `Update()` e `Draw()` próprios.

## Passo 4 – Sistema de cores
1. Defina cores fixas para cada time:
   - Time A: Verde
   - Time B: Azul
   - Bots: Vermelho
2. Use constantes ou `map[string]color.Color` para manter consistência.

## Passo 5 – Desenhar entidades simples
1. No `Draw`, desenhe formas básicas:
   - Jogador: quadrado verde ou azul
   - Bot: quadrado vermelho
   - Tiro: círculo branco ou amarelo
2. Use `ebitenutil.DrawRect` e `ebitenutil.DrawCircle` para protótipo rápido.

## Resultado esperado
- As entidades existem como tipos Go claros.
- O jogo pode atualizar e desenhar `Player`, `Bot` e `Projectile`.
- A arquitetura está organizada para incluir mais lógica.

## Próximo passo
Com as entidades definidas, o próximo objetivo é tornar o jogo concorrente usando goroutines e canais.
