package main

import (
	"image/color"
	"log"
	"math/rand"
	"time"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

const (
	screenWidth  = 640
	screenHeight = 480
	spriteSize   = 20 // Tamanho dos quadrados (jogadores e bots)
)

// Entidade representa qualquer objeto móvel no jogo (Jogador ou Bot)
type Entidade struct {
	X, Y  float64
	Color color.Color
}

// Jogo representa o estado principal do jogo
type Jogo struct {
	jogador *Entidade
	bots    []*Entidade
}

// NovoJogo inicializa o jogo com o jogador e bots
func NovoJogo() *Jogo {
	jogo_atual := &Jogo{
		jogador: &Entidade{X: screenWidth / 2, Y: screenHeight / 2, Color: color.RGBA{0, 255, 0, 255}}, // Verde
		bots:    make([]*Entidade, 0),
	}

	// Criar 2 bots iniciais
	jogo_atual.criarBot(100, 100)
	jogo_atual.criarBot(500, 400)

	return jogo_atual
}

// criarBot cria um novo bot e dispara sua Goroutine de IA
func (jogo_atual *Jogo) criarBot(x, y float64) {
	bot := &Entidade{X: x, Y: y, Color: color.RGBA{255, 0, 0, 255}} // Vermelho
	jogo_atual.bots = append(jogo_atual.bots, bot)

	// --- AQUI COMEÇA A CONCORRÊNCIA DO GO ---
	// Dispara uma Goroutine independente para controlar este bot específico
	go jogo_atual.menteDoBot(bot)
}

// menteDoBot roda em paralelo ao loop principal do jogo.
func (jogo_atual *Jogo) menteDoBot(b *Entidade) {
	// Ticker define a "velocidade de decisão" do bot (a cada 500ms)
	ticker := time.NewTicker(500 * time.Millisecond)
	defer ticker.Stop()

	// Seed para números aleatórios (para movimento errático)
	r := rand.New(rand.NewSource(time.Now().UnixNano()))

	for {
		select {
		case <-ticker.C:
			// Lógica simples: mover aleatoriamente
			// Em um sistema real, aqui você calcularia a perseguição ao jogador
			dx := float64(r.Intn(3) - 1) // -1, 0, ou 1
			dy := float64(r.Intn(3) - 1) // -1, 0, ou 1

			// Atualiza a posição (Nota: em produção, use Mutex aqui para segurança!)
			b.X += dx * spriteSize
			b.Y += dy * spriteSize

			// Impede o bot de sair da tela
			if b.X < 0 {
				b.X = 0
			}
			if b.X > screenWidth-spriteSize {
				b.X = screenWidth - spriteSize
			}
			if b.Y < 0 {
				b.Y = 0
			}
			if b.Y > screenHeight-spriteSize {
				b.Y = screenHeight - spriteSize
			}

			// case <-ctx.Done(): // Futuramente, usar context para matar o bot
			// 	return
		}
	}
}

// Update gerencia a lógica do jogo (principalmente input do usuário)
func (jogo_atual *Jogo) Update() error {
	// Controle do Jogador Principal (Setas do Teclado)
	speed := 4.0
	if ebiten.IsKeyPressed(ebiten.KeyLeft) {
		jogo_atual.jogador.X -= speed
	}
	if ebiten.IsKeyPressed(ebiten.KeyRight) {
		jogo_atual.jogador.X += speed
	}
	if ebiten.IsKeyPressed(ebiten.KeyUp) {
		jogo_atual.jogador.Y -= speed
	}
	if ebiten.IsKeyPressed(ebiten.KeyDown) {
		jogo_atual.jogador.Y += speed
	}

	// Impede o jogador de sair da tela
	if jogo_atual.jogador.X < 0 {
		jogo_atual.jogador.X = 0
	}
	if jogo_atual.jogador.X > screenWidth-spriteSize {
		jogo_atual.jogador.X = screenWidth - spriteSize
	}
	if jogo_atual.jogador.Y < 0 {
		jogo_atual.jogador.Y = 0
	}
	if jogo_atual.jogador.Y > screenHeight-spriteSize {
		jogo_atual.jogador.Y = screenHeight - spriteSize
	}

	return nil
}

// Draw desenha tudo na tela a cada frame (60 FPS)
func (jogo_atual *Jogo) Draw(screen *ebiten.Image) {
	// Desenha o fundo
	screen.Fill(color.RGBA{20, 20, 30, 255}) // Azul escuro

	// Desenha o Jogador (Quadrado Verde)
	ebitenutil.DrawRect(screen, jogo_atual.jogador.X, jogo_atual.jogador.Y, spriteSize, spriteSize, jogo_atual.jogador.Color)

	// Desenha os Bots (Quadrados Vermelhos)
	// Como os bots se movem em paralelo, o Draw apenas lê a posição atual deles
	for _, bot := range jogo_atual.bots {
		ebitenutil.DrawRect(screen, bot.X, bot.Y, spriteSize, spriteSize, bot.Color)
	}

	// Exibe instruções simples
	ebitenutil.DebugPrint(screen, "Use as SETAS para mover o Quadrado VERDE.\nOs VERMELHOS movem-se via Goroutines.")
}

// Layout define o tamanho lógico da tela
func (jogo_atual *Jogo) Layout(outsideWidth, outsideHeight int) (int, int) {
	return screenWidth, screenHeight
}

func main() {
	ebiten.SetWindowSize(screenWidth, screenHeight)
	ebiten.SetWindowTitle("Gopher Dungeon Arena - Protótipo Concorrente")

	game := NovoJogo()

	// Inicia o loop principal do jogo
	if err := ebiten.RunGame(game); err != nil {
		log.Fatal(err)
	}
}
