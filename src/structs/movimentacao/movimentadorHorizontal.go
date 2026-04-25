package movimentacao

import (
	"Gopher_Dungeon_Arena/src/configuracoes"
	"math/rand"
)

type MovimentadorHorizontal struct {
}

func (mb *MovimentadorHorizontal) Mover(objeto MovimentacaoCapacidade, r *rand.Rand) {
	posX := 0.0

	dX := float64(r.Intn(20) - 10)

	posX = objeto.GetX() + dX

	if posX >= configuracoes.PosXmax(10) {
		posX = configuracoes.PosXmax(10)
	} else if posX <= configuracoes.PosXmin() {
		posX = configuracoes.PosXmin()
	}

	objeto.SetPosicao(posX, objeto.GetY())
}
