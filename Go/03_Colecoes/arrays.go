package goguide

// Arrays em Go têm tamanho fixo e são úteis quando você conhece
// o número exato de elementos antecipadamente.
func FixedArrayExample() [3]int {
	arr := [3]int{1, 2, 3}
	return arr
}

// SumArray percorre um array fixo para somar seus elementos.
func SumArray(values [3]int) int {
	total := 0
	for _, v := range values {
		total += v
	}
	return total
}

// UpdateArray mostra como modificar elementos em um array pelo índice.
func UpdateArray(values [3]int) [3]int {
	values[0] = 10
	return values
}
