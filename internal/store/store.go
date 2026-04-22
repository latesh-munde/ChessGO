package store

import "chess-server/internal/game"

type GameStore interface {
	Save(*game.Game)
	Get(id string) (*game.Game, error)
	Join(id string) error
}
