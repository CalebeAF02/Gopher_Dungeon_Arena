package goguide

import "fmt"

// A função init é executada automaticamente antes de main.
// Ela é útil para inicializar estados ou configurar o pacote.
func init() {
	fmt.Println("init() executado antes do main")
}
