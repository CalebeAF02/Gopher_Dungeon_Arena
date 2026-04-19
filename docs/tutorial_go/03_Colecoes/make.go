package goguide

// MakeSlice mostra como pré-alocar um slice com tamanho zero e capacidade definida.
func MakeSlice(capacity int) []int {
	return make([]int, 0, capacity)
}

// MakeMap demonstra criação de um map com capacidade inicial.
func MakeMap(capacity int) map[string]int {
	return make(map[string]int, capacity)
}

// SliceWithLength cria um slice com tamanho e capacidade iguais.
func SliceWithLength(length int) []int {
	return make([]int, length)
}
