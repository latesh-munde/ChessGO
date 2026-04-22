package pieces

import "chess-server/internal/core"

type PieceType int

const (
	King PieceType = iota
	Queen
	Rook
	Bishop
	Knight
	Pawn
)

type Piece interface {
	Type() PieceType
	Color() core.Color
	LegalMoves(from core.Square, board core.Board) []core.Square
}

type BasePiece struct {
	color core.Color
}

func (b BasePiece) Color() core.Color {
	return b.color
}

func (p PieceType) String() string {
	switch p {
	case King:
		return "King"
	case Queen:
		return "Queen"
	case Rook:
		return "Rook"
	case Bishop:
		return "Bishop"
	case Knight:
		return "Knight"
	case Pawn:
		return "Pawn"
	default:
		return "Unknown"
	}
}
