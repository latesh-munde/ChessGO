package pieces

import "chess-server/internal/core"

type PawnPiece struct {
	BasePiece
}

func NewPawn(c core.Color) *PawnPiece {
	return &PawnPiece{BasePiece{c}}
}

func (p *PawnPiece) Type() PieceType {
	return Pawn
}

func (p *PawnPiece) LegalMoves(from core.Square, board core.Board) []core.Square {
	var moves []core.Square

	dir := 1
	startRank := 1
	if p.color == core.Black {
		dir = -1
		startRank = 6
	}

	// forward
	if sq, ok := from.Offset(dir, 0); ok && board.Empty(sq) {
		moves = append(moves, sq)

		if from.Rank == startRank {
			if sq2, ok := from.Offset(dir*2, 0); ok && board.Empty(sq2) {
				moves = append(moves, sq2)
			}
		}
	}

	// captures
	for _, dc := range []int{-1, 1} {
		if sq, ok := from.Offset(dir, dc); ok && board.EnemyAt(sq, p.color) {
			moves = append(moves, sq)
		}
	}

	return moves
}
