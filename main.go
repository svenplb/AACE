package main

import (
	"fmt"

	"github.com/notnil/chess"
)

func main() {
	game := chess.NewGame()
	// print outcome and game PGN
	moves := game.ValidMoves()

	fmt.Println(game.Position().Board().Draw())
	fmt.Println(moves)

	if err := game.MoveStr("e4"); err != nil {
		// handle error
		fmt.Println("es funktioniert garnichts bro, was machst du")
	}

	fmt.Println(game.String())
	fmt.Println(game.Position().Board().Draw())
	fmt.Println(moves)

	if err := game.MoveStr("e5"); err != nil {
		// handle error
		fmt.Println("es funktioniert garnichts bro, was machst du")
	}

	// history
	fmt.Println(game.String())
	// draw board
	fmt.Println(game.Position().Board().Draw())
	// valid moves
	fmt.Println(moves)

}
