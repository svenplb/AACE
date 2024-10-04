package handlers

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/notnil/chess"
)

func alphaBetaSearch_old(position *chess.Position, game *chess.Game, depth int, alpha int, beta int) (int, *chess.Move) {
	if depth == 0 || position.Status() != chess.NoMethod {
		return Evaluate_old(position), nil
	}
	moves := position.ValidMoves()
	if len(moves) == 0 {
		return Evaluate_old(position), nil
	}
	var bestMove *chess.Move

	if position.Turn() == chess.White {
		bestEval := -9999
		for _, move := range moves {
			newPosition := position.Update(move)
			eval, _ := alphaBetaSearch_old(newPosition, game, depth-1, alpha, beta)
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
			eval, _ := alphaBetaSearch_old(newPosition, game, depth-1, alpha, beta)
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

func Search_old(c *gin.Context) {

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

	eval, bestMove := alphaBetaSearch_old(game.Position(), game, 3, -9999, 9999)

	fmt.Println(bestMove)

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
