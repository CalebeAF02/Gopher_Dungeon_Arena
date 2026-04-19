package goguide

// O tipo booleano em Go é bool e pode ser true ou false.
// Ele é usado em expressões condicionais e operadores lógicos.

func BooleanOperators(a, b bool) (and, or, notA bool) {
	// and: verdadeiro se ambos forem verdadeiros.
	and = a && b
	// or: verdadeiro se ao menos um for verdadeiro.
	or = a || b
	// notA: nega o valor de a.
	notA = !a
	return
}

// IsAdult verifica se uma idade indica maioridade.
// Operadores relacionais retornam bools.
func IsAdult(age int) bool {
	return age >= 18
}

// BooleanFromComparison demonstra como o resultado de uma comparação
// é automaticamente um valor booleano.
func BooleanFromComparison(a, b int) bool {
	return a > b
}
