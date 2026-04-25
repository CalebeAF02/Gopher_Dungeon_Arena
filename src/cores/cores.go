package cores

import (
	"image/color"
)

var (
	AMARELO      = COR(255, 235, 59)
	VERMELHO     = COR(229, 57, 53)
	AZUL         = COR(3, 155, 229)
	LARANJA      = COR(255, 87, 34)
	VERDE        = COR(76, 175, 80)
	VERDE_ESCURO = COR(56, 142, 60)
	TURQUESA     = COR(69, 39, 160)
	BRANCO       = COR(227, 242, 253)
)

func COR(r uint8, g uint8, b uint8) color.Color {
	return color.RGBA{r, g, b, 255}
}
