package engine

import (
	"chess-server/internal/core"
	"chess-server/internal/game"
	"chess-server/internal/pieces"
)

func CanCastle(state *game.State, color core.Color, kingSide bool) bool {
	rank := 0
	if color == core.Black {
		rank = 7
	}

	kingFile := 4
	rookFile := 7
	step := 1
	if !kingSide {
		rookFile = 0
		step = -1
	}

	king := state.Board.Squares[rank][kingFile]
	rook := state.Board.Squares[rank][rookFile]

	if king == nil || rook == nil {
		return false
	}

	if king.Type() != pieces.King || rook.Type() != pieces.Rook {
		return false
	}

	for f := kingFile + step; f != rookFile; f += step {
		if state.Board.Squares[rank][f] != nil {
			return false
		}
	}

	return true
}
