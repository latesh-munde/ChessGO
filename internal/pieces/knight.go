package pieces

import "chess-server/internal/core"

type KnightPiece struct {
	BasePiece
}

func NewKnight(c core.Color) *KnightPiece {
	return &KnightPiece{BasePiece{c}}
}

func (n *KnightPiece) Type() PieceType {
	return Knight
}

func (n *KnightPiece) LegalMoves(from core.Square, board core.Board) []core.Square {
	dirs := []struct{ r, c int }{
		{2, 1}, {2, -1}, {-2, 1}, {-2, -1},
		{1, 2}, {1, -2}, {-1, 2}, {-1, -2},
	}

	var moves []core.Square
	for _, d := range dirs {
		if sq, ok := from.Offset(d.r, d.c); ok {
			if board.Empty(sq) || board.EnemyAt(sq, n.color) {
				moves = append(moves, sq)
			}
		}
	}
	return moves
}
