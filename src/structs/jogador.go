package structs

import (
	"Gopher_Dungeon_Arena/src/ecs"
	"image/color"

	"fmt"
)

type Jogador struct {
	game     *Game
	entidade ecs.Entidade
	nome     string
	vida     int
	sangue   int
	cor      color.Color
	status   bool
	vivendo  int
}

func NovoJogador(game *Game, n string) *Jogador {
	nEntidade := game.CriarEntidade()

	nJogador := Jogador{game: game, entidade: nEntidade, nome: n, vida: 2, sangue: 100, cor: color.White, status: true, vivendo: 0}
	game.Posicoes[nEntidade] = ecs.NovaPosicao(0, 0)

	game.Jogadores[nEntidade] = &nJogador

	return &nJogador
}

func (j *Jogador) estaVivo() bool {
	if j.vida == 0 {
		j.status = false
	}
	return j.status
}

func (j *Jogador) renasce() {
	if j.sangue == 0 {
		j.vida -= 1
		if j.estaVivo() {
			fmt.Println("O jogador " + j.nome + " morreu!")
		}
		j.resetaSangue()
	}
}

func (j *Jogador) resetaSangue() {
	if j.vida == 1 {
		j.sangue = 100
	}
}

func (j *Jogador) GetNome() string {
	return j.nome
}

func (j *Jogador) SetPosicao(x float64, y float64) {
	j.game.Posicoes[j.entidade].SetPosicao(x, y)
}

func (j *Jogador) GetX() float64 {
	return j.game.Posicoes[j.entidade].GetX()
}

func (j *Jogador) GetY() float64 {
	return j.game.Posicoes[j.entidade].GetY()
}

func (j *Jogador) GetCor() color.Color {
	return j.cor
}

func (j *Jogador) SetCor(cor color.Color) {
	j.cor = cor
}

func (j *Jogador) GetVivendo() int {
	return j.vivendo
}

func (j *Jogador) Atualizar() {
	j.vivendo += 1
	if j.vivendo >= 30 {
		j.vivendo = 0
	}
}
