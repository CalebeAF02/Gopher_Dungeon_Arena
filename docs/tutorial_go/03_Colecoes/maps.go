package goguide

// CreateMap demonstra criação de um map e adição de elementos.
func CreateMap() map[string]int {
	m := map[string]int{"maçã": 3, "banana": 5}
	m["laranja"] = 2
	return m
}

// GetValueFromMap mostra como ler um valor e verificar se a chave existe.
func GetValueFromMap(m map[string]int, key string) (int, bool) {
	value, ok := m[key]
	return value, ok
}

// DeleteFromMap remove uma chave do map usando delete.
func DeleteFromMap(m map[string]int, key string) {
	delete(m, key)
}

// MapSize retorna quantos pares chave-valor existem no map.
func MapSize(m map[string]int) int {
	return len(m)
}
