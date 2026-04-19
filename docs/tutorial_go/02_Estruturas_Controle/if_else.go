package goguide

// CheckSign retorna uma descrição com base no valor numérico.
// if, else if e else são usados para tomar decisões em tempo de execução.
func CheckSign(value int) string {
	if value > 0 {
		return "positivo"
	} else if value < 0 {
		return "negativo"
	}
	return "zero"
}

// IsEven mostra que qualquer expressão booleana pode ser usada em um if.
func IsEven(value int) bool {
	if value%2 == 0 {
		return true
	}
	return false
}

// CanVote demonstra uma condicional composta com comparação de idade.
func CanVote(age int) bool {
	if age >= 16 && age < 18 {
		return true
	} else if age >= 18 {
		return true
	}
	return false
}
