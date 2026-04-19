package structs

import "fmt"

type Jogador struct {
	nome   string
	vida   int
	sangue int
	status bool
}

func NovoJogador(n string) Jogador {
	return Jogador{nome: n, vida: 2, sangue: 100, status: true}
}

func (jogador *Jogador) estaVivo() bool {
	if jogador.vida == 0 {
		jogador.status = false
	}
	return jogador.status
}

func (jogador *Jogador) renasce() {
	if jogador.sangue == 0 {
		jogador.vida -= 1
		if jogador.estaVivo() {
			fmt.Println("O jogador " + jogador.nome + " morreu!")
		}
		jogador.resetaSangue()
	}
}

func (jogador *Jogador) resetaSangue() {
	if jogador.vida == 1 {
		jogador.sangue = 100
	}
}

func (jogador *Jogador) GetNome() string {
	return jogador.nome
}
