package pieces

import "chess-server/internal/core"

type RookPiece struct {
	BasePiece
}

func NewRook(c core.Color) *RookPiece {
	return &RookPiece{BasePiece{c}}
}

func (r *RookPiece) Type() PieceType {
	return Rook
}

func (r *RookPiece) LegalMoves(from core.Square, board core.Board) []core.Square {
	return slidingMoves(
		from,
		board,
		r.color,
		[]struct{ r, c int }{
			{1, 0}, {-1, 0}, {0, 1}, {0, -1},
		},
	)
}
