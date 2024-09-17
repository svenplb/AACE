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

// search algorithm, return best evaluation and move
func searchTree(position *chess.Position, depth int, alpha, beta int, maximizingPlayer bool) (int, *chess.Move) {
	if depth == 0 {
		return evaluateBoardPosition(position), nil
	}

	moves := position.ValidMoves()
	var bestMove *chess.Move

	if maximizingPlayer {
		maxEval := -9999
		for _, move := range moves {
			newPosition := position.Update(move)
			eval, _ := searchTree(newPosition, depth-1, alpha, beta, false)
			if eval > maxEval {
				maxEval = eval
				bestMove = move
			}
			alpha = max(alpha, eval)
			if beta <= alpha {
				break
			}
		}
		return maxEval, bestMove
	} else {
		minEval := 9999
		for _, move := range moves {
			newPosition := position.Update(move)
			eval, _ := searchTree(newPosition, depth-1, alpha, beta, true)
			if eval < minEval {
				minEval = eval
				bestMove = move
			}
			beta = min(beta, eval)
			if beta <= alpha {
				break
			}
		}
		return minEval, bestMove
	}
}

// Helper functions
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func playFullGame(game *chess.Game) {

	for game.Outcome() == chess.NoOutcome {
		moves := game.ValidMoves()
		if len(moves) == 0 {
			break
		}
		bestEvaluation, bestMove := searchTree(game.Position(), 6, -9999, 9999, true)
		fmt.Println("Best Move: " + bestMove.String())
		fmt.Println("Best Evaluation: ", bestEvaluation)
		err := game.Move(bestMove)

		if err != nil {
			fmt.Println("Error making move:", err)
			break
		}

		fmt.Println(game.Position().Board().Draw())

	}

}

func main() {
	fenStr := "rn2k1r1/ppp1pp1p/3p2p1/5bn1/P7/2N2B2/1PPPPP2/2BNK1RR w Gkq - 4 11"
	fen, _ := chess.FEN(fenStr)
	game := chess.NewGame(fen)
	playFullGame(game)

}
