package main

import (
	"Gopher_Dungeon_Arena/src/configuracoes"
	"Gopher_Dungeon_Arena/src/structs"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
)

func main() {
	game := structs.NovoGame()
	ebiten.SetWindowSize(configuracoes.LARGURA, configuracoes.ALTURA)
	ebiten.SetWindowTitle(configuracoes.NOME_JOGO)
	if err := ebiten.RunGame(&game); err != nil {
		log.Fatal(err)
	}
}
