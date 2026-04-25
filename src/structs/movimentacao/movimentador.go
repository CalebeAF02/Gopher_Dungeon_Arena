package movimentacao

import (
	"math/rand"
)

type MovimentacaoCapacidade interface {
	GetX() float64
	GetY() float64
	SetPosicao(x, y float64)
}

type Movimentador interface {
	Mover(bot MovimentacaoCapacidade, r *rand.Rand)
}
