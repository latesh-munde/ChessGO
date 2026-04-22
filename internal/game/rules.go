package game

import (
	"chess-server/internal/core"
	"chess-server/internal/pieces"
)

func ApplyMove(s *State, m Move) {
	p := s.Board.Squares[m.From.Rank][m.From.File]

	// clear source
	s.Board.Squares[m.From.Rank][m.From.File] = nil

	// en-passant target
	s.EnPassant = nil
	if p.Type() == pieces.Pawn {
		if abs(m.From.Rank-m.To.Rank) == 2 {
			mid := (m.From.Rank + m.To.Rank) / 2
			s.EnPassant = &core.Square{
				Rank: mid,
				File: m.From.File,
			}
		}
	}

	// place piece
	s.Board.Squares[m.To.Rank][m.To.File] = p

	// switch turn
	s.Turn = s.Turn.Opponent()
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
