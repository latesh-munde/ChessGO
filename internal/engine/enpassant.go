package engine

import (
	"chess-server/internal/core"
	"chess-server/internal/game"
)

// HandleEnPassant removes captured pawn if en-passant is played
func HandleEnPassant(s *game.State, m game.Move) {
	if s.EnPassant == nil {
		return
	}

	// If move lands on en-passant square
	if m.To == *s.EnPassant {
		dir := -1
		if s.Turn == core.White {
			dir = 1
		}

		captured := core.Square{
			Rank: m.To.Rank + dir,
			File: m.To.File,
		}

		s.Board.Squares[captured.Rank][captured.File] = nil
	}

	// Clear en-passant after any move
	s.EnPassant = nil
}
