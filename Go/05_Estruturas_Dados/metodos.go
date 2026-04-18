package goguide

import "fmt"

// Enemy representa um inimigo em um jogo.
type Enemy struct {
	Name  string
	Power int
}

// Description é um método de Enemy que retorna uma descrição textual.
func (e Enemy) Description() string {
	return fmt.Sprintf("%s tem poder de %d", e.Name, e.Power)
}

// BoostPower demonstra um método que altera o estado do receiver.
func (e *Enemy) BoostPower(amount int) {
	e.Power += amount
}
