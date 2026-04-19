package structs


type Time struct {
	Nome      string
	Jogadores []Jogador
	Vivo      bool
}

func (t *Time) AdicionnarJogador(jogador Jogador) {
	t.Jogadores = append(t.Jogadores, jogador)
}

