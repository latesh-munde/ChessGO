package engine

import (
	"chess-server/internal/core"
	"chess-server/internal/game"
)

func HasAnyLegalMove(s *game.State, color core.Color) bool {
	for r := 0; r < 8; r++ {
		for f := 0; f < 8; f++ {
			p := s.Board.Squares[r][f]
			if p == nil || p.Color() != color {
				continue
			}

			from := core.Square{Rank: r, File: f}
			for _, to := range p.LegalMoves(from, s.Board) {
				m := game.Move{From: from, To: to}
				if !WouldLeaveKingInCheck(s, m) {
					return true
				}
			}
		}
	}
	return false
}

func IsCheckmate(s *game.State, color core.Color) bool {
	return InCheck(s, color) && !HasAnyLegalMove(s, color)
}

func IsStalemate(s *game.State, color core.Color) bool {
	return !InCheck(s, color) && !HasAnyLegalMove(s, color)
}
