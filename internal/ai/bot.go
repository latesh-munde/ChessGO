package ai

import "chess-server/internal/game"

// public interface
type Bot interface {
	GetMove(state *game.State) (game.Move, error)
}
