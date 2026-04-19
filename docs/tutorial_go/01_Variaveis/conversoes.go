package goguide

import (
	"fmt"
	"strconv"
)

// ConverteStringParaInt converte uma string para inteiro usando strconv.Atoi.
// A função retorna o valor convertido e o erro, que deve ser tratado pelo chamador.
func ConverteStringParaInt(valor string) (int, error) {
	numero, err := strconv.Atoi(valor)
	if err != nil {
		return 0, fmt.Errorf("não foi possível converter %q para int: %w", valor, err)
	}
	return numero, nil
}

// ConverteIntParaString converte um inteiro para string usando strconv.Itoa.
func ConverteIntParaString(numero int) string {
	return strconv.Itoa(numero)
}

// ConverteStringParaBool converte uma string para bool usando strconv.ParseBool.
// A string deve ser "1", "t", "T", "true", "TRUE", "True", "0", "f", "F",
// "false", "FALSE" ou "False".
func ConverteStringParaBool(valor string) (bool, error) {
	return strconv.ParseBool(valor)
}

// ConverteBoolParaString converte um booleano para string.
func ConverteBoolParaString(valor bool) string {
	return strconv.FormatBool(valor)
}

// ConversoesExemplo demonstra conversões de string para int, int para string,
// bool para string e string para bool.
func ConversoesExemplo() string {
	numero, err := ConverteStringParaInt("42")
	if err != nil {
		return err.Error()
	}
	texto := ConverteIntParaString(numero)
	booleano, err := ConverteStringParaBool("true")
	if err != nil {
		return err.Error()
	}
	return fmt.Sprintf("numero=%d texto=%q booleano=%t", numero, texto, booleano)
}
