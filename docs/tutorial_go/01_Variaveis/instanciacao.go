package goguide

import "fmt"

// Instanciação de variáveis em Go pode ser feita de várias maneiras.
// 1) Com a palavra-chave var e tipo explícito.
// 2) Com var e inferência de tipo automática.
// 3) Com := onde o tipo é inferido a partir do valor.
// 4) Com new() para obter um ponteiro para o tipo.

// Declaração explícita com tipo e valor inicial.
var DefaultName string = "Gopher"

// Declaração explícita com tipo sem valor inicial atribuído. Recebe o valor zero padrão.
var Number int

func ExampleDeclarations() string {
	// Inferência de tipo: o compilador determina o tipo pelo valor literal.
	nome := "Go"
	idade := 30

	// Tipo explícito sem valor inicial: recebe zero value (0 para inteiros, "" para strings, false para bool).
	var ativo bool

	return fmt.Sprintf("%s tem %d anos e ativo=%v", nome, idade, ativo)
}

// NewPointer mostra como usar new() para criar uma variável alocada no heap e obter seu ponteiro.
func NewPointer() *int {
	p := new(int)
	*p = 100
	return p
}

// ZeroValues demonstra o valor padrão de tipos básicos em Go.
func ZeroValues() (int, float64, string, bool) {
	var inteiro int
	var flutuante float64
	var texto string
	var booleano bool
	return inteiro, flutuante, texto, booleano
}

// Constantes não são variáveis, mas são importantes para comparar com variáveis
// porque não podem ser modificadas após a inicialização.
const Pi float64 = 3.1415926535
