package handlers

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

// algorithm is kind of broken
func alphaBeta(position *chess.Position, depth int, alpha int, beta int, isMaximizingPlayer bool) (int, *chess.Move) {
	if depth == 0 || position.Status() != chess.NoMethod {
		return evaluateBoardPosition(position), nil
	}

	moves := position.ValidMoves()
	if len(moves) == 0 {
		return evaluateBoardPosition(position), nil
	}

	var bestMove *chess.Move
	if isMaximizingPlayer {
		maxEval := -99999
		for _, move := range moves {
			newPosition := position.Update(move)
			eval, _ := alphaBeta(newPosition, depth-1, alpha, beta, false)
			if eval > maxEval {
				maxEval = eval
				bestMove = move
			}
			alpha = max(alpha, eval)
			if beta <= alpha {
				break // Beta cut-off
			}
		}
		return maxEval, bestMove
	} else {
		minEval := 99999
		for _, move := range moves {
			newPosition := position.Update(move)
			eval, _ := alphaBeta(newPosition, depth-1, alpha, beta, true)
			if eval < minEval {
				minEval = eval
				bestMove = move
			}
			beta = min(beta, eval)
			if beta <= alpha {
				break // Alpha cut-off
			}
		}
		return minEval, bestMove
	}
}

func maxx(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func minn(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func engineAgainstEngine(game *chess.Game) {
	alpha := -99999
	beta := 99999

	for game.Outcome() == chess.NoOutcome {
		moves := game.ValidMoves()
		if len(moves) == 0 {
			break
		}

		if game.Position().Turn() == chess.White {
			fmt.Println("White is thinking")
		} else {
			fmt.Println("Black is thinking")
		}

		bestEvaluation, bestMove := alphaBeta(game.Position(), 7, alpha, beta, game.Position().Turn() == chess.White)

		fmt.Println("Best Move:", bestMove, "Eval of Move:", bestEvaluation)
		err := game.Move(bestMove)

		if err != nil {
			fmt.Println("Error making move:", err)
			break
		}

		fmt.Println(game.Position().Board().Draw())
	}
}

func humanAgainstEngine(game *chess.Game, depth int) {
	alpha := -99999
	beta := 99999

	fmt.Println(game.Position().Board().Draw()) // Display initial board

	for game.Outcome() == chess.NoOutcome {

		if game.Position().Turn() == chess.White { // Human (White)
			fmt.Println("Your turn (White). Enter move:")

			validMoves := game.Position().ValidMoves()
			fmt.Println("Valid moves:", validMoves)

			var move string
			fmt.Scan(&move)

			if err := game.MoveStr(move); err != nil {
				fmt.Println("INVALID MOVE:", move)
				fmt.Println("Error:", err)
				continue
			}

		} else {
			fmt.Println("Anti Adrian Chessbot (Black) is thinking...")

			bestEvaluation, bestMove := alphaBeta(game.Position(), depth, alpha, beta, game.Position().Turn() == chess.White)

			fmt.Println("Engine found best move:", bestMove, "with evaluation:", bestEvaluation)

			err := game.Move(bestMove)
			if err != nil {
				fmt.Println("Error making move:", err)
				break
			}
		}

		fmt.Println(game.Position().Board().Draw())
	}

	switch game.Outcome() {
	case chess.WhiteWon:
		fmt.Println("You won! Congratulations!")
	case chess.BlackWon:
		fmt.Println("The engine won. L!")
	case chess.Draw:
		fmt.Println("It's a draw!")
	default:
		fmt.Println("Game over.")
	}
}

func singlePositionSearch(game *chess.Game) {
	alpha := -99999
	beta := -99999
	bestEvaluation, bestMove := alphaBeta(game.Position(), 6, alpha, beta, game.Position().Turn() == chess.White)
	fmt.Println("Best Eval", bestEvaluation, "best Move found:", bestMove)

}

func main() {
	fenStr := "r1bk3r/p2pBpNp/n4n2/1p1NP2P/6P1/3P4/P1P1K3/q5b1"
	fen, _ := chess.FEN(fenStr)
	game := chess.NewGame(fen)
	fmt.Println(game.Position().Turn())
	//singlePositionSearch(game)
	engineAgainstEngine(game)
	//humanAgainstEngine(game, 6)

}
