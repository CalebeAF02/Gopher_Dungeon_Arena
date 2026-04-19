package goguide

// SumAll soma um número variável de argumentos inteiros.
func SumAll(values ...int) int {
	total := 0
	for _, v := range values {
		total += v
	}
	return total
}

// JoinWords concatena um número variável de palavras separadas por espaço.
func JoinWords(words ...string) string {
	result := ""
	for i, word := range words {
		if i > 0 {
			result += " "
		}
		result += word
	}
	return result
}
