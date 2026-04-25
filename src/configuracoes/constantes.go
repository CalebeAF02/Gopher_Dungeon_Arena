package configuracoes

import (
	"math/rand"
	"time"
)

const (
	LARGURA   = 800
	ALTURA    = 600
	NOME_JOGO = "Gopher_Dungeon_Arena"
)

func PosXmax(largura_obj int) float64 {
	return float64(LARGURA - largura_obj)
}

func PosXmin() float64 {
	return 0.0
}

func PosYmax(altura_obj int) float64 {
	return float64(ALTURA - altura_obj)
}

func PosYmin() float64 {
	return 0.0
}

func XAleatorio(r *rand.Rand) float64 {
	return float64(r.Intn(LARGURA))
}

func YAleatorio(r *rand.Rand) float64 {
	return float64(r.Intn(ALTURA))
}

func GeradorAleatorio() *rand.Rand {
	return rand.New(rand.NewSource(time.Now().UnixNano()))
}
