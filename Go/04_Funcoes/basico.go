package goguide

// Add é um exemplo de função básica com dois parâmetros e retorno de valor.
func Add(a int, b int) int {
	return a + b
}

// Greet mostra como receber uma string e retornar outra string formatada.
func Greet(name string) string {
	return "Olá, " + name
}

// SplitFullName demonstra retorno de múltiplos valores separados por vírgula.
func SplitFullName(fullName string) (string, string) {
	// Em um exemplo real, poderíamos separar o nome e sobrenome.
	return fullName, ""
}
