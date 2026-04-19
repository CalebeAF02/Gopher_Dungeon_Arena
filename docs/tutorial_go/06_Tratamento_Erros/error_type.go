package goguide

import "errors"

// NewSimpleError mostra como criar um erro simples usando errors.New.
func NewSimpleError() error {
	return errors.New("algo deu errado")
}

// FormatError cria uma mensagem de erro formatada com informações adicionais.
func FormatError(field string) error {
	return errors.New("campo inválido: " + field)
}
