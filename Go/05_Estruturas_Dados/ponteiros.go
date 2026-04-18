package goguide

// Increment demonstra como um ponteiro permite alterar o valor original.
func Increment(value *int) {
	*value++
}

// NewPointerExample retorna um ponteiro para um valor inteiro alocado no stack.
func NewPointerExample() *int {
	x := 7
	return &x
}

// PointerNil mostra o zero value de um ponteiro, que é nil.
func PointerNil() *int {
	var p *int
	return p
}
