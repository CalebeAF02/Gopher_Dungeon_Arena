package goguide

// GoModInfo explica o propósito do arquivo go.mod.
// O go.mod define o nome do módulo, a versão do Go e as dependências externas.
// Ele também torna o projeto portável e fácil de construir em outras máquinas.
func GoModInfo() string {
	return "go.mod gerencia dependências e versão do módulo"
}

// ExampleModuleContent mostra como seria o conteúdo de um go.mod básico.
func ExampleModuleContent() string {
	return `module github.com/usuario/meu-projeto

go 1.22

require (
    github.com/gin-gonic/gin v1.10.0
)`
}
