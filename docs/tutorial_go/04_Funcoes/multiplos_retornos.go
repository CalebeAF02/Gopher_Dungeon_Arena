package goguide

import "errors"

// Divide mostra como retornar um valor e um erro quando a operação pode falhar.
func Divide(a, b float64) (float64, error) {
	if b == 0 {
		return 0, errors.New("divisão por zero")
	}
	return a / b, nil
}

// CurrencyChange calcula o troco e retorna também se houve erro.
func CurrencyChange(amount, price float64) (float64, error) {
	if amount < price {
		return 0, errors.New("valor insuficiente")
	}
	return amount - price, nil
}
