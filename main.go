package main

import (
	"log"
	"tictac/game"
	"github.com/hajimehoshi/ebiten/v2"
)

func main() {
	game := game.NewGame()

	ebiten.SetWindowSize(300, 300)
	ebiten.SetWindowTitle("Jogo da Velha")

	if err := ebiten.RunGame(game); err != nil {
		log.Fatal(err)
	}
}
