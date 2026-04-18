package goguide

import "fmt"

// Drawable define um contrato que exige o método Draw.
type Drawable interface {
	Draw() string
}

// Circle representa uma forma circular.
type Circle struct {
	Radius int
}

// Draw implementa a interface Drawable para Circle.
func (c Circle) Draw() string {
	return fmt.Sprintf("Desenhando círculo de raio %d", c.Radius)
}

// Rectangle representa uma forma retangular.
type Rectangle struct {
	Width  int
	Height int
}

// Draw implementa a interface Drawable para Rectangle.
func (r Rectangle) Draw() string {
	return fmt.Sprintf("Desenhando retângulo %dx%d", r.Width, r.Height)
}
