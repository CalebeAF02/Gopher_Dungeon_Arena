# Meta 1: Fundação e Game Loop

Este arquivo mostra os passos para montar a infraestrutura inicial do jogo em Go usando Ebiten.

## Passo 1 – Inicializar o módulo Go
1. Abra o terminal na pasta do projeto.
2. Execute:
   ```bash
   go mod init gopher-dungeon
   go get github.com/hajimehoshi/ebiten/v2
   ```
3. Verifique que `go.mod` e `go.sum` foram criados.

## Passo 2 – Criar o arquivo principal
1. Dentro de `src/`, abra ou crie `main.go`.
2. Defina o pacote e as importações:
   ```go
   package main

   import (
       "log"
       "github.com/hajimehoshi/ebiten/v2"
       "github.com/hajimehoshi/ebiten/v2/ebitenutil"
   )
   ```

## Passo 3 – Definir a estrutura do jogo
1. Crie uma struct `Game`:
   ```go
   type Game struct {}
   ```
2. Implemente os métodos obrigatórios do Ebiten:
   - `Update() error`
   - `Draw(screen *ebiten.Image)`
   - `Layout(outsideWidth, outsideHeight int) (int, int)`

## Passo 4 – Implementar o loop básico
1. Em `Update`, apenas retorne `nil` por enquanto.
2. Em `Draw`, desenhe um texto simples:
   ```go
   func (g *Game) Draw(screen *ebiten.Image) {
       ebitenutil.DebugPrint(screen, "Gopher Dungeon PvPvE")
   }
   ```
3. Em `Layout`, retorne a largura e altura da janela, por exemplo `800, 600`.

## Passo 5 – Executar o jogo
1. No `main()`, inicialize o jogo:
   ```go
   func main() {
       ebiten.SetWindowSize(800, 600)
       ebiten.SetWindowTitle("Gopher Dungeon PvPvE")
       if err := ebiten.RunGame(&Game{}); err != nil {
           log.Fatal(err)
       }
   }
   ```
2. Execute com:
   ```bash
   go run ./src
   ```
3. Confirme que a janela abre e exibe o texto.

## Resultado esperado
- Janela do jogo abre.
- Texto simples aparece na tela.
- O loop do Ebiten está funcionando.

## Próximo passo
Após confirmar essa base, avance para criar as entidades e o sistema de tipagem do jogo.
