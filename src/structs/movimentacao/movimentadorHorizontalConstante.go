package movimentacao

import (
	"Gopher_Dungeon_Arena/src/configuracoes"
	"math/rand"
)

type MovimentadorHorizontalConstante struct {
	dx     int
	ciclos int
}

func (mb *MovimentadorHorizontalConstante) Mover(objeto MovimentacaoCapacidade, r *rand.Rand) {

	if mb.dx == 0 {
		mb.dx = r.Intn(10)
		mb.ciclos = 0
	}

	posX := objeto.GetX() + float64(mb.dx)

	if posX >= configuracoes.PosXmax(10) {
		posX = configuracoes.PosXmax(10)

		mb.dx = r.Intn(10)
		mb.dx = mb.dx * (-1)

	} else if posX <= configuracoes.PosXmin() {
		posX = configuracoes.PosXmin()

		mb.dx = r.Intn(10)
		mb.dx = mb.dx * (-1)
	}

	mb.ciclos += 1

	if mb.ciclos == 10 {
		mb.dx = r.Intn(10)
	}

	objeto.SetPosicao(posX, objeto.GetY())
}
