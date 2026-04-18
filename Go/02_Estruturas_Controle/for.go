package goguide

// SumRange usa o loop clássico for com inicialização, condição e incremento.
func SumRange(n int) int {
	total := 0
	for i := 0; i < n; i++ {
		total += i
	}
	return total
}

// MultiplyItems demonstra o uso de for com índice para percorrer um slice.
func MultiplyItems(values []int) int {
	result := 1
	for i := 0; i < len(values); i++ {
		result *= values[i]
	}
	return result
}

// CountdownSteps usa um for para reduzir o valor até zero.
func CountdownSteps(start int) []int {
	steps := []int{}
	for i := start; i > 0; i-- {
		steps = append(steps, i)
	}
	return steps
}
