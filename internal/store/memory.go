package store

import (
	"errors"
	"sync"

	"chess-server/internal/game"
)

type MemoryStore struct {
	mu    sync.RWMutex
	games map[string]*game.Game
}

func NewMemoryStore() *MemoryStore {
	return &MemoryStore{games: make(map[string]*game.Game)}
}

func (s *MemoryStore) Save(g *game.Game) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.games[g.ID] = g
}

func (s *MemoryStore) Get(id string) (*game.Game, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()
	g, ok := s.games[id]
	if !ok {
		return nil, errors.New("not found")
	}
	return g, nil
}

func (s *MemoryStore) Join(id string) error {
	_, err := s.Get(id)
	return err
}
