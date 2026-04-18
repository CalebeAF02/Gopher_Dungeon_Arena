package goguide

// CreateSlice demonstra um slice dinâmico com append.
func CreateSlice() []int {
	s := []int{10, 20, 30}
	s = append(s, 40)
	return s
}

// CopySlice cria uma cópia independente de um slice.
func CopySlice(source []int) []int {
	dest := make([]int, len(source))
	copy(dest, source)
	return dest
}

// SubSlice mostra como criar uma fatia a partir de um slice existente.
func SubSlice(source []int) []int {
	if len(source) < 3 {
		return source
	}
	return source[1:3]
}

// CapacityExample retorna o comprimento e a capacidade de um slice.
func CapacityExample() (int, int) {
	s := make([]int, 2, 5)
	return len(s), cap(s)
}
