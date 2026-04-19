package main

import (
	"Golpher_Dungeon_Arena/src/structs"
	"fmt"
)

func main() {
	t1 := structs.Time{}
	j1 := structs.NovoJogador("Calebe")
	t1.AdicionnarJogador(j1)

	for _, jogador := range t1.Jogadores {
		fmt.Println("Jogador: " + jogador.GetNome())
	}
}
func loop() {
	atualiza()
	desenha()
}

func atualiza() {

}

func desenha() {

}
