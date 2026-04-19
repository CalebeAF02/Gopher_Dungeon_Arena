package goguide

// Em Go não existe uma palavra reservada while.
// O mesmo comportamento é obtido usando for com uma condição.
func CountDown(start int) []int {
	result := []int{}
	for start > 0 {
		result = append(result, start)
		start--
	}
	return result
}

// FindFirstEven mostra como implementar um loop while usando for.
func FindFirstEven(numbers []int) int {
	i := 0
	for i < len(numbers) {
		if numbers[i]%2 == 0 {
			return numbers[i]
		}
		i++
	}
	return -1
}

// WhileExample demonstra um loop que continua até uma condição ser falsa.
func WhileExample(limit int) int {
	total := 0
	i := 0
	for i < limit {
		total += i
		i++
	}
	return total
}
