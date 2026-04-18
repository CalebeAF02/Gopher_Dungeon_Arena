package goguide

// DayOfWeek demonstra o uso de switch para avaliar vários casos.
func DayOfWeek(day int) string {
	switch day {
	case 1:
		return "Domingo"
	case 2:
		return "Segunda"
	case 3:
		return "Terça"
	case 4:
		return "Quarta"
	case 5:
		return "Quinta"
	case 6:
		return "Sexta"
	case 7:
		return "Sábado"
	default:
		return "Desconhecido"
	}
}

// DescribeType mostra um exemplo de type switch para descobrir o tipo real.
func DescribeType(value interface{}) string {
	switch value.(type) {
	case int:
		return "inteiro"
	case string:
		return "texto"
	case bool:
		return "booleano"
	case float64:
		return "float64"
	default:
		return "tipo desconhecido"
	}
}

// FeelsLikeWeekend é um exemplo com múltiplos valores no mesmo case.
func FeelsLikeWeekend(day int) string {
	switch day {
	case 6, 7:
		return "fim de semana"
	default:
		return "dia de semana"
	}
}
