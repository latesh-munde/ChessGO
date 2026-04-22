package engine

import (
	"errors"

	"chess-server/internal/game"
)

// ValidateMove validates a move fully (rules + king safety)
func ValidateMove(s *game.State, m game.Move) error {
	p := s.Board.Squares[m.From.Rank][m.From.File]
	if p == nil {
		return errors.New("no piece on source square")
	}

	if p.Color() != s.Turn {
		return errors.New("not your turn")
	}

	legal := p.LegalMoves(m.From, s.Board)
	ok := false
	for _, sq := range legal {
		if sq == m.To {
			ok = true
			break
		}
	}
	if !ok {
		return errors.New("illegal move")
	}

	if WouldLeaveKingInCheck(s, m) {
		return errors.New("move leaves king in check")
	}

	return nil
}
