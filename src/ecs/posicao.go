package ecs

type Posicao struct {
	posX, posY float64
}

func NovaPosicao(x float64, y float64) *Posicao {
	return &Posicao{posX: x, posY: y}
}

func (p *Posicao) GetX() float64 {
	return p.posX
}

func (p *Posicao) GetY() float64 {
	return p.posY
}

func (p *Posicao) SetX(x float64) {
	p.posX = x
}

func (p *Posicao) SetY(y float64) {
	p.posY = y
}

func (p *Posicao) SetPosicao(x float64, y float64) {
	p.posX = x
	p.posY = y
}
