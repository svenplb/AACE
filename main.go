package main

import (
	"fmt"

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

// func searchTree(position *chess.Position, depth int, alpha int, beta int) int {}
func searchTree(position *chess.Position, depth int) (int, *chess.Move) {
	// depth = how deep the alg / tree goes
	// alpha = best score for maximizing
	// beta = best score for  minimizing
	moves := position.ValidMoves()

	// return current state if depth == 0
	if depth == 0 {
		return evaluateBoardPosition(position), nil
	}

	bestEvaluation := -9999
	var bestMove *chess.Move
	for _, move := range moves {
		newPosition := position.Update(move)
		evaluation, _ := searchTree(newPosition, depth-1)
		if evaluation > bestEvaluation {
			bestEvaluation = evaluation
			bestMove = move
		}

	}
	return bestEvaluation, bestMove

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
		bestEvaluation, bestMove := searchTree(game.Position(), 3)
		fmt.Println("Best Move: " + bestMove.String())
		fmt.Println("Best Evaluation: ", bestEvaluation)
		err := game.Move(bestMove)

		if err != nil {
			fmt.Println("Error making move:", err)
			break
		}

		score := evaluateBoardPosition(game.Position())
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

		fmt.Println("Game outcome:", game.Outcome())
		fmt.Println(game.Position().Board().Draw())
		fmt.Println(moves)

	}

}
