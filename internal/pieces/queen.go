package pieces

import "chess-server/internal/core"

type QueenPiece struct {
	BasePiece
}

func NewQueen(c core.Color) *QueenPiece {
	return &QueenPiece{BasePiece{c}}
}

func (q *QueenPiece) Type() PieceType {
	return Queen
}

func (q *QueenPiece) LegalMoves(from core.Square, board core.Board) []core.Square {
	return slidingMoves(
		from,
		board,
		q.color,
		[]struct{ r, c int }{
			{1, 0}, {-1, 0}, {0, 1}, {0, -1},
			{1, 1}, {1, -1}, {-1, 1}, {-1, -1},
		},
	)
}
