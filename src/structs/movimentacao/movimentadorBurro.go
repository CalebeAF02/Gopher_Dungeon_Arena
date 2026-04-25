package movimentacao

import (
	"Gopher_Dungeon_Arena/src/configuracoes"
	"math/rand"
)

type MovimentadorBurro struct {
}

func (mb *MovimentadorBurro) Mover(objeto MovimentacaoCapacidade, r *rand.Rand) {
	posX := 0.0
	posY := 0.0

	dX := float64(r.Intn(10) - 5)
	dY := float64(r.Intn(10) - 5)

	posX = objeto.GetX() + dX
	posY = objeto.GetY() + dY

	if posX >= configuracoes.PosXmax(10) {
		posX = configuracoes.PosXmax(10)
	} else if posX <= configuracoes.PosXmin() {
		posX = configuracoes.PosXmin()
	}

	if posY <= configuracoes.PosYmin() {
		posY = configuracoes.PosYmin()
	} else if posY >= configuracoes.PosYmax(10) {
		posY = configuracoes.PosYmax(10)
	}

	objeto.SetPosicao(posX, posY)

}
