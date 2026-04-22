package pieces

import "chess-server/internal/core"

type KingPiece struct {
	BasePiece
}

func NewKing(c core.Color) *KingPiece {
	return &KingPiece{BasePiece{c}}
}

func (k *KingPiece) Type() PieceType {
	return King
}

func (k *KingPiece) LegalMoves(from core.Square, board core.Board) []core.Square {
	dirs := []struct{ r, c int }{
		{1, 0}, {-1, 0}, {0, 1}, {0, -1},
		{1, 1}, {1, -1}, {-1, 1}, {-1, -1},
	}

	var moves []core.Square
	for _, d := range dirs {
		if sq, ok := from.Offset(d.r, d.c); ok {
			if board.Empty(sq) || board.EnemyAt(sq, k.color) {
				moves = append(moves, sq)
			}
		}
	}
	return moves
}
