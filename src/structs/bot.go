package structs

import (
	"Gopher_Dungeon_Arena/src/ecs"
	"Gopher_Dungeon_Arena/src/structs/movimentacao"
	"image/color"
	"math/rand"
)

type Bot struct {
	game     *Game
	entidade ecs.Entidade
	Id       int64
	sangue   int
	cor      color.Color
	status   bool
	movendo  movimentacao.Movimentador
}

func NovoBot(game *Game, id int64) *Bot {
	verde := color.RGBA{118, 255, 3, 255}

	nEntidade := game.CriarEntidade()

	nBot := Bot{game: game, entidade: nEntidade, Id: id, sangue: 100, cor: verde, status: true}
	game.Posicoes[nEntidade] = ecs.NovaPosicao(0, 0)

	game.Bots[nEntidade] = &nBot

	return &nBot
}

func (b *Bot) perseguir() {

}

func (b *Bot) SetPosicao(x float64, y float64) {
	b.game.Posicoes[b.entidade].SetPosicao(x, y)
}

func (b *Bot) GetX() float64 {
	return b.game.Posicoes[b.entidade].GetX()
}

func (b *Bot) GetY() float64 {
	return b.game.Posicoes[b.entidade].GetY()
}

func (b *Bot) GetCor() color.Color {
	return b.cor
}

func (b *Bot) SetCor(c color.Color) {
	b.cor = c
}

func (b *Bot) Mover(r *rand.Rand) {
	if b.movendo != nil {
		b.movendo.Mover(b, r)
	}
}

func (b *Bot) SetMovimentacao(movendo movimentacao.Movimentador) {
	b.movendo = movendo
}
