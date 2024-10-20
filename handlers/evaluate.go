package handlers

import (
	"github.com/notnil/chess"
)

func Evaluate(position *chess.Position) int {

	var pieceValues = map[chess.PieceType]int{
		chess.King:   999999,
		chess.Queen:  900,
		chess.Rook:   500,
		chess.Bishop: 300,
		chess.Knight: 300,
		chess.Pawn:   100,
	}

	pawn_w_sv := [64]int{
		0, 0, 0, 0, 0, 0, 0, 0,
		5, 10, 10, -20, -20, 10, 10, 5,
		5, -5, -10, 0, 0, -10, -5, 5,
		0, 0, 0, 20, 20, 0, 0, 0,
		5, 5, 10, 25, 25, 10, 5, 5,
		10, 10, 20, 30, 30, 20, 10, 10,
		50, 50, 50, 50, 50, 50, 50, 50,
		0, 0, 0, 0, 0, 0, 0, 0,
	}

	knight_w_sv := [64]int{
		-50, -40, -30, -30, -30, -30, -40, -50,
		-40, -20, 0, 5, 5, 0, -20, -40,
		-30, 5, 10, 15, 15, 10, 5, -30,
		-30, 0, 15, 20, 20, 15, 0, -30,
		-30, 5, 15, 20, 20, 15, 5, -30,
		-30, 0, 10, 15, 15, 10, 0, -30,
		-40, -20, 0, 0, 0, 0, -20, -40,
		-50, -40, -30, -30, -30, -30, -40, -50,
	}

	bishop_w_sv := [64]int{
		-20, -10, -10, -10, -10, -10, -10, -20,
		-10, 5, 0, 0, 0, 0, 5, -10,
		-10, 10, 10, 10, 10, 10, 10, -10,
		-10, 0, 10, 10, 10, 10, 0, -10,
		-10, 5, 5, 10, 10, 5, 5, -10,
		-10, 0, 5, 10, 10, 5, 0, -10,
		-10, 0, 0, 0, 0, 0, 0, -10,
		-20, -10, -10, -10, -10, -10, -10, -20,
	}

	rook_w_sv := [64]int{
		0, 0, 5, 10, 10, 5, 0, 0,
		-5, 0, 0, 0, 0, 0, 0, -5,
		-5, 0, 0, 0, 0, 0, 0, -5,
		-5, 0, 0, 0, 0, 0, 0, -5,
		-5, 0, 0, 0, 0, 0, 0, -5,
		-5, 0, 0, 0, 0, 0, 0, -5,
		5, 10, 10, 10, 10, 10, 10, 5,
		0, 0, 0, 0, 0, 0, 0, 0,
	}

	queen_w_sv := [64]int{
		-20, -10, -10, -5, -5, -10, -10, -20,
		-10, 0, 5, 0, 0, 0, 0, -10,
		-10, 5, 5, 5, 5, 5, 0, -10,
		0, 0, 5, 5, 5, 5, 0, -5,
		-5, 0, 5, 5, 5, 5, 0, -5,
		-10, 0, 5, 5, 5, 5, 0, -10,
		-10, 0, 0, 0, 0, 0, 0, -10,
		-20, -10, -10, -5, -5, -10, -10, -20,
	}

	king_w_sv := [64]int{
		20, 30, 10, 0, 0, 10, 30, 20,
		20, 20, 0, 0, 0, 0, 20, 20,
		-10, -20, -20, -20, -20, -20, -20, -10,
		-20, -30, -30, -40, -40, -30, -30, -20,
		-30, -40, -40, -50, -50, -40, -40, -30,
		-30, -40, -40, -50, -50, -40, -40, -30,
		-30, -40, -40, -50, -50, -40, -40, -30,
		-30, -40, -40, -50, -50, -40, -40, -30,
	}

	var pawn_b_sv [64]int
	var knight_b_sv [64]int
	var bishop_b_sv [64]int
	var rook_b_sv [64]int
	var queen_b_sv [64]int
	var king_b_sv [64]int

	// ! Wtf is this shit, fix this ASAP. Clunky ahh code
	for i := 0; i < 64; i++ {
		pawn_b_sv[i] = pawn_w_sv[63-i]
		knight_b_sv[i] = knight_w_sv[63-i]
		bishop_b_sv[i] = bishop_w_sv[63-i]
		rook_b_sv[i] = rook_w_sv[63-i]
		queen_b_sv[i] = queen_w_sv[63-i]
		king_b_sv[i] = king_w_sv[63-i]
	}

	score := 0
	board := position.Board()

	for sq := 0; sq < 64; sq++ {
		piece := board.Piece(chess.Square(sq))
		if piece != chess.NoPiece {
			value := pieceValues[piece.Type()]
			if piece.Color() == chess.White {
				score += value
				switch piece.Type() {
				case chess.Pawn:
					score += pawn_w_sv[sq]
				case chess.Knight:
					score += knight_w_sv[sq]
				case chess.Bishop:
					score += bishop_w_sv[sq]
				case chess.Rook:
					score += rook_w_sv[sq]
				case chess.Queen:
					score += queen_w_sv[sq]
				case chess.King:
					score += king_w_sv[sq]
				}
			} else {
				score -= value
				switch piece.Type() {
				case chess.Pawn:
					score -= pawn_b_sv[sq]
				case chess.Knight:
					score -= knight_b_sv[sq]
				case chess.Bishop:
					score -= bishop_b_sv[sq]
				case chess.Rook:
					score -= rook_b_sv[sq]
				case chess.Queen:
					score -= queen_b_sv[sq]
				case chess.King:
					score -= king_b_sv[sq]
				}

			}
		}

	}

	friendlyKingSquare := -1
	opponentKingSquare := -1
	for sq := 0; sq < 64; sq++ {
		piece := board.Piece(chess.Square(sq))
		if piece.Type() == chess.King {
			if piece.Color() == chess.White {
				friendlyKingSquare = sq
			} else {
				opponentKingSquare = sq
			}
		}
	}
	if friendlyKingSquare != -1 && opponentKingSquare != -1 {
		score += ForceKingToCornerEndgameEval(friendlyKingSquare, opponentKingSquare, 5.5)
	}

	return score

}

// Helper functions for endgame

func Rank(square int) int {
	return square / 8
}

func File(square int) int {
	return square % 8
}

func Abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// TODO: Play around with values.
// eval position of both kings. Try to push opponent king to corner & keep friendly king safe.
func ForceKingToCornerEndgameEval(friendlyKingSquare int, opponentKingSquare int, endgameWeight float64) int {
	evaluation := 0

	// row / column opponent king
	opponentKingRank := Rank(opponentKingSquare)
	opponentKingFile := File(opponentKingSquare)

	// distance to center row / colum & add up distances
	opponentKingDstToCentreFile := Max(3-opponentKingFile, opponentKingFile-4)
	opponentKingDstToCentreRank := Max(3-opponentKingRank, opponentKingRank-4)

	opponentKingDstFromCentre := opponentKingDstToCentreFile + opponentKingDstToCentreRank
	evaluation += opponentKingDstFromCentre

	friendlyKingRank := Rank(friendlyKingSquare)
	friendlyKingFile := File(friendlyKingSquare)

	dstBetweenKingsFile := Abs(friendlyKingFile - opponentKingFile)
	dstBetweenKingsRank := Abs(friendlyKingRank - opponentKingRank)

	dstBetweenKings := dstBetweenKingsFile + dstBetweenKingsRank
	evaluation += 14 - dstBetweenKings

	return int(float64(evaluation) * 10 * endgameWeight)
}
