package game

import (
	"chess-server/internal/core"
	"chess-server/internal/pieces"
)

type Board struct {
	Squares [8][8]pieces.Piece
}

func NewBoard() *Board {
	return &Board{}
}

func (b *Board) Empty(s core.Square) bool {
	return b.Squares[s.Rank][s.File] == nil
}

func (b *Board) EnemyAt(s core.Square, c core.Color) bool {
	p := b.Squares[s.Rank][s.File]
	return p != nil && p.Color() != c
}

func (b *Board) FriendlyAt(s core.Square, c core.Color) bool {
	p := b.Squares[s.Rank][s.File]
	return p != nil && p.Color() == c
}
