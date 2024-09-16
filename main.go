package main

import (
	"fmt"
	"math/rand"

	"github.com/notnil/chess"
)

var pieceValues = map[chess.PieceType]int{
	chess.Queen:  9,
	chess.Rook:   5,
	chess.Bishop: 3,
	chess.Knight: 3,
	chess.Pawn:   1,
}

func evaluateBoardPosition(position *chess.Position) int {
	score := 0
	board := position.Board()

	for sq := 0; sq < 64; sq++ {
		piece := board.Piece(chess.Square(sq))
		if piece != chess.NoPiece {
			value := pieceValues[piece.Type()]
			if piece.Color() == chess.White {
				score += value
			} else {
				score -= value
			}

		}
	}
	return score
}

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
		score := (evaluateBoardPosition(game.Position()))
		if score > 0 {
			fmt.Println("White is winning")
			fmt.Println(score)
		} else if score < 0 {
			fmt.Println("Black is winning")
			fmt.Println(score)
		} else {
			fmt.Println("The game is drawn")
			fmt.Println(score)
		}
		fmt.Println(game.Position().Board().Draw())
	}

}
