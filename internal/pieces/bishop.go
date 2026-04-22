package pieces

import "chess-server/internal/core"

type BishopPiece struct {
	BasePiece
}

func NewBishop(c core.Color) *BishopPiece {
	return &BishopPiece{BasePiece{c}}
}

func (b *BishopPiece) Type() PieceType {
	return Bishop
}

func (b *BishopPiece) LegalMoves(from core.Square, board core.Board) []core.Square {
	return slidingMoves(
		from,
		board,
		b.color,
		[]struct{ r, c int }{
			{1, 1}, {1, -1}, {-1, 1}, {-1, -1},
		},
	)
}
