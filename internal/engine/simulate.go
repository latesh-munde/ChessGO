package engine

import (
	"chess-server/internal/game"
)

// WouldLeaveKingInCheck simulates a move and checks king safety
func WouldLeaveKingInCheck(s *game.State, m game.Move) bool {
	from := m.From
	to := m.To

	p := s.Board.Squares[from.Rank][from.File]
	captured := s.Board.Squares[to.Rank][to.File]

	// simulate
	s.Board.Squares[from.Rank][from.File] = nil
	s.Board.Squares[to.Rank][to.File] = p

	inCheck := InCheck(s, p.Color())

	// rollback
	s.Board.Squares[from.Rank][from.File] = p
	s.Board.Squares[to.Rank][to.File] = captured

	return inCheck
}
