package goguide

import "fmt"

// ValidationError é um tipo de erro customizado que contém campo e mensagem.
type ValidationError struct {
	Field   string
	Message string
}

// Error implementa a interface error para ValidationError.
func (e ValidationError) Error() string {
	return fmt.Sprintf("campo %s: %s", e.Field, e.Message)
}

// NewValidationError cria um ValidationError com informações específicas.
func NewValidationError(field, message string) error {
	return ValidationError{Field: field, Message: message}
}

// IsValidationError verifica o tipo do erro usando type assertion.
func IsValidationError(err error) bool {
	_, ok := err.(ValidationError)
	return ok
}
