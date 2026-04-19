package goguide

import "fmt"

// OperadoresAritmeticos demonstra as operações básicas de soma, subtração,
// multiplicação, divisão e resto em Go.
func OperadoresAritmeticos(a, b int) string {
	soma := a + b
	subtracao := a - b
	multiplicacao := a * b
	divisao := a / b
	resto := a % b
	return fmt.Sprintf("soma=%d subtracao=%d multiplicacao=%d divisao=%d resto=%d", soma, subtracao, multiplicacao, divisao, resto)
}

// OperadoresAtribuicao mostra como atualizar a variável utilizando operadores
// como +=, -=, *= e /=.
func OperadoresAtribuicao(x int) string {
	x += 5
	x -= 2
	x *= 3
	x /= 2
	return fmt.Sprintf("resultado=%d", x)
}

// OperadoresRelacionais compara valores e retorna resultados booleanos.
func OperadoresRelacionais(a, b int) (igual, diferente, maior, menor bool) {
	igual = a == b
	diferente = a != b
	maior = a > b
	menor = a < b
	return
}

// ConcatenaStrings demonstra concatenação de strings com +.
func ConcatenaStrings(a, b string) string {
	return a + " " + b
}

// InterpolacaoStrings demonstra a interpolação de strings com fmt.Sprintf.
func InterpolacaoStrings(nome string, idade int) string {
	return fmt.Sprintf("%s tem %d anos", nome, idade)
}
