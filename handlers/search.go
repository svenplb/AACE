package handlers

import (
	"fmt"
	"net/http"
	"sort"

	"github.com/gin-gonic/gin"
	"github.com/notnil/chess"
)

func alphaBetaSearch(position *chess.Position, game *chess.Game, depth int, alpha int, beta int, searchCountt *int) (int, *chess.Move) {
	*searchCountt++
	if depth == 0 || position.Status() != chess.NoMethod {
		return Evaluate(position), nil
	}

	moves := orderMoves(position.ValidMoves(), position)

	if len(moves) == 0 {
		return Evaluate(position), nil
	}
	var bestMove *chess.Move

	if position.Turn() == chess.White {
		bestEval := -9999
		for _, move := range moves {

			newPosition := position.Update(move)
			eval, _ := alphaBetaSearch(newPosition, game, depth-1, alpha, beta, searchCountt)
			if eval > bestEval {
				bestEval = eval
				bestMove = move
			}
			if eval > alpha {
				alpha = eval
			}
			if beta <= alpha {
				break // Beta
			}
		}
		return bestEval, bestMove
	} else {
		bestEval := 9999
		for _, move := range moves {
			newPosition := position.Update(move)
			eval, _ := alphaBetaSearch(newPosition, game, depth-1, alpha, beta, searchCountt)
			if eval < bestEval {
				bestEval = eval
				bestMove = move
			}
			if eval < beta {
				beta = eval
			}

			if beta <= alpha {
				break // Alpha
			}
		}
		return bestEval, bestMove
	}
}

func orderMoves(validMoves []*chess.Move, position *chess.Position) []*chess.Move {
	type scoredMove struct {
		move  *chess.Move
		score int
	}

	scoredMoves := make([]scoredMove, 0, len(validMoves))

	for _, move := range validMoves {
		movePiece := position.Board().Piece(move.S1())

		capturePiece := position.Board().Piece(move.S2())

		moveScore := 0

		switch movePiece.Type() {
		case chess.Queen:
			moveScore = 9
		case chess.Rook:
			moveScore = 5
		case chess.Bishop, chess.Knight:
			moveScore = 3
		case chess.Pawn:
			moveScore = 1
		case chess.King:
			moveScore = 10 //
		}

		if capturePiece != chess.NoPiece {
			switch capturePiece.Type() {
			case chess.Queen:
				moveScore += 9
			case chess.Rook:
				moveScore += 5
			case chess.Bishop, chess.Knight:
				moveScore += 3
			case chess.Pawn:
				moveScore += 1
			}
		}

		scoredMoves = append(scoredMoves, scoredMove{move: move, score: moveScore})
	}

	sort.Slice(scoredMoves, func(i, j int) bool {
		return scoredMoves[i].score > scoredMoves[j].score
	})

	sortedMoves := make([]*chess.Move, len(scoredMoves))
	for i, sm := range scoredMoves {
		sortedMoves[i] = sm.move
	}

	return sortedMoves
}
func Search(c *gin.Context) {
	var json struct {
		FEN string `json:"fen"`
	}

	if err := c.BindJSON(&json); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	fen, err := chess.FEN(json.FEN)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error:": "Invalid FEN"})
		return
	}

	game := chess.NewGame(fen)

	searchCountt := 0

	eval, bestMove := alphaBetaSearch(game.Position(), game, 6, -9999, 9999, &searchCountt)

	fmt.Printf("Positions searched (MOVE ORDERING): %d\n", searchCountt)

	if bestMove == nil {
		c.JSON(http.StatusOK, gin.H{"error": "No valid moves found"})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"best_move":      bestMove.String(),
			"evaluation":     eval,
			"fen_after_move": game.FEN(),
		})
	}

}
