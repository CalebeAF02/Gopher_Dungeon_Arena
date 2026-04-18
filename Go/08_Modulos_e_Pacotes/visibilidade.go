package goguide

// Tipos e funções com letra maiúscula são exportados e podem ser usados por outros pacotes.
// Tipos e funções com letra minúscula são privados ao pacote atual.

type privateStruct struct {
	Field string
}

type ExportedStruct struct {
	PublicField  string
	privateField string
}

// ExportedFunction é visível fora do pacote goguide.
func ExportedFunction() string {
	return "função exportada"
}

// unexportedFunction só pode ser usada dentro do pacote goguide.
func unexportedFunction() string {
	return "função não exportada"
}

// NewExportedStruct cria um ExportedStruct inicializado.
func NewExportedStruct(public string, private string) ExportedStruct {
	return ExportedStruct{PublicField: public, privateField: private}
}
