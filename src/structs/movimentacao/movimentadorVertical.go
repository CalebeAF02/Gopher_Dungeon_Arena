package movimentacao

import (
	"Gopher_Dungeon_Arena/src/configuracoes"
	"math/rand"
)

type MovimentadorVertical struct {
}

func (mb *MovimentadorVertical) Mover(objeto MovimentacaoCapacidade, r *rand.Rand) {
	posY := 0.0

	dY := float64(r.Intn(12) - 6)

	posY = objeto.GetY() + dY

	if posY >= configuracoes.PosYmax(10) {
		posY = configuracoes.PosYmax(10)
	} else if posY <= configuracoes.PosYmin() {
		posY = configuracoes.PosYmin()
	}

	objeto.SetPosicao(objeto.GetX(), posY)
}
