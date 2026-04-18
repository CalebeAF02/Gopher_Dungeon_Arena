package goguide

// Tipos complexos em Go são complex64 e complex128.
// Cada valor complex contém uma parte real e uma parte imaginária.

func ComplexValues() (complex64, complex128) {
	var a complex64 = 2 + 3i
	var b complex128 = complex(1.5, -4.25)
	return a, b
}

// ComplexOperations demonstra operações básicas com números complexos.
func ComplexOperations() complex128 {
	a := complex(1.0, 2.0)
	b := complex(3.0, 4.0)
	return a + b
}
