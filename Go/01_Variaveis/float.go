package goguide

// Tipos de ponto flutuante em Go são float32 e float64.
// Eles representam números com parte fracionária, mas não são exatos para todos
// os valores decimais devido à forma como são armazenados em binário.

func FloatPrecision() (float32, float64) {
	var x float32 = 1.2345678
	var y float64 = 1.23456789012345
	return x, y
}

// AddFloats soma dois valores de ponto flutuante float32.
func AddFloats(a float32, b float32) float32 {
	return a + b
}

// FloatComparison mostra que comparações diretas podem ser perigosas
// quando há pequenas diferenças de precisão.
func FloatComparison(a, b float64) bool {
	const epsilon = 1e-9
	diff := a - b
	if diff < 0 {
		diff = -diff
	}
	return diff < epsilon
}
