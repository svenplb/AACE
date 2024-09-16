package main

import (
	"fmt"
	"math/rand"

	"github.com/notnil/chess"
)

func main() {
	game := chess.NewGame()
	// print outcome and game PGN
	moves := game.ValidMoves()

	fmt.Println(game.Position().Board().Draw())
	fmt.Println(moves)

	for game.Outcome() == chess.NoOutcome {
		moves := game.ValidMoves()
		if len(moves) == 0 {
			break
		}

		randomMove := moves[rand.Intn(len(moves))]

		err := game.Move(randomMove)
		if err != nil {
			fmt.Println("Error making move:", err)
			break

		}
		fmt.Printf("Move made: %s\n", randomMove.String())
		fmt.Println(game.Position().Board().Draw())
		fmt.Println(game.String())
	}

}
