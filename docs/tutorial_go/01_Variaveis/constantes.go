package goguide

import "fmt"

// ConstantesGo demonstra declaração de constantes em Go.
const (
	VersaoDaLinguagem = "1.22"
	StatusAtivo       = true
)

// ConstantesNumericas demonstra o uso de iota para gerar valores sequenciais.
const (
	Primeiro = iota // 0
	Segundo         // 1
	Terceiro        // 2
)

// ExemploConstantes retorna os valores das constantes declaradas.
func ExemploConstantes() string {
	return fmt.Sprintf("versao=%s ativo=%t primeiro=%d segundo=%d terceiro=%d", VersaoDaLinguagem, StatusAtivo, Primeiro, Segundo, Terceiro)
}
