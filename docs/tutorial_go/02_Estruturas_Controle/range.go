package goguide

// RangeSlice demonstra iteração em um slice usando range.
func RangeSlice(values []int) int {
	total := 0
	for _, v := range values {
		total += v
	}
	return total
}

// RangeMap mostra como iterar sobre um map, recebendo chave e valor.
func RangeMap(data map[string]int) []string {
	keys := make([]string, 0, len(data))
	for k, _ := range data {
		keys = append(keys, k)
	}
	return keys
}

// RangeValues retorna todos os valores de um map em uma lista.
func RangeValues(data map[string]int) []int {
	values := make([]int, 0, len(data))
	for _, v := range data {
		values = append(values, v)
	}
	return values
}
