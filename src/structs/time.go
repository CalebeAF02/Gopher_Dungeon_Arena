package structs

import (
	"image/color"
)

type Time struct {
	nome      string
	jogadores []*Jogador
	cor      color.Color
}

func NovoTime(game *Game, n string, cor color.Color) Time {
	t := Time{nome: n, cor: cor}
	return t
}

func (t *Time) Adicionnar(jogador *Jogador) {
	jogador.SetCor(t.cor)
	t.jogadores = append(t.jogadores, jogador)
}

func (t *Time) EstaVivo() bool {
	for _, jogador := range t.jogadores {
		if jogador.status {
			return true
		}
	}
	return false
}

func (t *Time) GetNome() string {
	return t.nome
}

func (t *Time) GetJogadores() []*Jogador {
	return t.jogadores
}
