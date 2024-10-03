package handlers

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/notnil/chess"
)

// ! DIESRE ALOGIRHTMUS IS FUCKED, KEINE AHUNNT WIE dA IRGENDWAS FUNCKTIONIERT.
// TODO: Rewrite that bih
func minimax(position *chess.Position, depth int) (int, *chess.Move) {
	if depth == 0 || position.Status() != chess.NoMethod {
		return evaluateBoardPosition(position), nil
	}
	moves := position.ValidMoves()
	if len(moves) == 0 {
		return evaluateBoardPosition(position), nil
	}
	var bestMove *chess.Move
	var bestEval int
	if position.Turn() == chess.White {
		bestEval = -9999
		for _, move := range moves {
			fmt.Println(move.String())
			newPosition := position.Update(move)
			eval, _ := minimax(newPosition, depth-1)
			if eval > bestEval {
				bestEval = eval
				bestMove = move
			}
		}
	} else {
		bestEval = 9999
		for _, move := range moves {
			fmt.Println(move.String())
			newPosition := position.Update(move)
			eval, _ := minimax(newPosition, depth-1)
			if eval < bestEval {
				bestEval = eval
				bestMove = move
			}
		}
	}
	return bestEval, bestMove
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

	_, bestMove := minimax(game.Position(), 3)

	fmt.Println(bestMove)
	eval := Evaluate(game)

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
