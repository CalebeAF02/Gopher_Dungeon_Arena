package goguide

import "fmt"

// Strings em Go são sequências imutáveis de bytes UTF-8.
// Elas podem ser criadas com aspas duplas ou com crases para literais brutos.

// ConcatStrings junta duas strings usando o operador +.
func ConcatStrings(a, b string) string {
	return a + b
}

// Interpolate monta uma mensagem usando fmt.Sprintf.
// Este é o jeito mais comum de inserir valores em uma string.
func Interpolate(name string, age int) string {
	return fmt.Sprintf("Meu nome é %s e tenho %d anos.", name, age)
}

// RawStringLiteral demonstra uma string de múltiplas linhas
// usando crases (`), preservando quebras e espaços exatamente.
func RawStringLiteral() string {
	return `Esta é uma string
multilinha em Go.
Ela mantém todos os caracteres literais.`
}

// StringLength retorna o comprimento da string em bytes.
// No caso de caracteres UTF-8 com mais de um byte, o resultado
// é o tamanho em bytes, não em letras visuais.
func StringLength(text string) int {
	return len(text)
}
