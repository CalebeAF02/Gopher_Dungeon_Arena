package structs

import (
	"Gopher_Dungeon_Arena/src/configuracoes"
	"Gopher_Dungeon_Arena/src/cores"
	"Gopher_Dungeon_Arena/src/ecs"
	"Gopher_Dungeon_Arena/src/structs/movimentacao"
	"fmt"
	"image/color"
	"math/rand"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

type Game struct {
	proximo   int
	Jogadores map[ecs.Entidade]*Jogador
	Bots      map[ecs.Entidade]*Bot
	Posicoes  map[ecs.Entidade]*ecs.Posicao
	aleatorio *rand.Rand

	time1 Time
	time2 Time
}

func (g *Game) Update() error {

	for _, bot := range g.Bots {
		bot.Mover(g.aleatorio)
	}

	for _, jogador := range g.Jogadores {
		jogador.Atualizar()
	}

	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	screen.Fill(color.RGBA{20, 20, 20, 255})
	// Desenhar Time 1 (Red)
	for _, jogador := range g.Jogadores {
		ebitenutil.DrawRect(screen, jogador.GetX(), jogador.GetY(), 20, 20, jogador.GetCor())

		ebitenutil.DrawRect(screen, jogador.GetX()+5, jogador.GetY()+5, 5, 5, color.White)
		ebitenutil.DrawRect(screen, jogador.GetX()+10, jogador.GetY()+5, 5, 5, color.White)

		ebitenutil.DrawRect(screen, jogador.GetX()+5, jogador.GetY()+10, 5, 5, color.White)
		ebitenutil.DrawRect(screen, jogador.GetX()+10, jogador.GetY()+10, 5, 5, color.White)

	}
	for _, bot := range g.Bots {

		ebitenutil.DrawRect(screen, bot.GetX(), bot.GetY(), 10, 10, bot.GetCor())
	}
}

func (g *Game) Layout(l, a int) (int, int) {
	return configuracoes.LARGURA, configuracoes.ALTURA
}

func NovoGame() Game {
	jogadores := make(map[ecs.Entidade]*Jogador)
	bots := make(map[ecs.Entidade]*Bot)
	posicoes := make(map[ecs.Entidade]*ecs.Posicao)
	aleatorio := configuracoes.GeradorAleatorio()

	g := Game{Jogadores: jogadores, Bots: bots, Posicoes: posicoes, aleatorio: aleatorio}

	// Jogadores
	j1 := NovoJogador(&g, "Jogador 1")
	j2 := NovoJogador(&g, "Jogador 2")
	j3 := NovoJogador(&g, "Jogador 3")

	j4 := NovoJogador(&g, "Jogador 4")
	j5 := NovoJogador(&g, "Jogador 5")
	j6 := NovoJogador(&g, "Jogador 6")

	for id := 0; id < 20; id++ {
		b := NovoBot(&g, int64(id))
		b.SetPosicao(configuracoes.XAleatorio(aleatorio), configuracoes.YAleatorio(aleatorio))

		if aleatorio.Intn(100) < 30 {
			b.SetMovimentacao(&movimentacao.MovimentadorBurro{})
			b.SetCor(cores.TURQUESA)
		} else {
			valor :=  aleatorio.Intn(100)
			if valor< 30 {
				b.SetMovimentacao(&movimentacao.MovimentadorVertical{})
				b.SetCor(cores.LARANJA)
			}else 			if valor > 30 && valor<50 {
				b.SetMovimentacao(&movimentacao.MovimentadorHorizontalConstante{})
				b.SetCor(cores.BRANCO)
			} else {
				b.SetMovimentacao(&movimentacao.MovimentadorHorizontal{})
				b.SetCor(cores.AZUL)

			}
		}

		fmt.Printf("x: %f | Y: %f\n", b.GetX(), b.GetY())
	}

	j1.SetPosicao(100, 500)
	j2.SetPosicao(200, 300)
	j3.SetPosicao(300, 500)

	j4.SetPosicao(500, 100)
	j5.SetPosicao(300, 200)
	j6.SetPosicao(500, 300)

	// Times
	t1 := NovoTime(&g, "Red", color.RGBA{255, 0, 0, 255})
	t2 := NovoTime(&g, "Blue", color.RGBA{0, 0, 255, 255})

	// Gerenciando
	t1.Adicionnar(j1)
	t1.Adicionnar(j2)
	t1.Adicionnar(j3)

	t2.Adicionnar(j4)
	t2.Adicionnar(j5)
	t2.Adicionnar(j6)

	g.SetTimes(t1, t2)
	return g
}

func (g *Game) CriarEntidade() ecs.Entidade {
	entidade := ecs.Entidade(g.proximo)
	g.proximo++
	return entidade
}

func (g *Game) SetTimes(t1 Time, t2 Time) {
	g.time1 = t1
	g.time2 = t2
}
