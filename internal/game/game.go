package game

import (
	"chess-server/internal/core"
	"errors"
	"sync"
	"time"

	"github.com/google/uuid"
)

type Game struct {
	ID    string
	State *State
	Moves []Move
	mu    sync.Mutex
	Ended bool
}

func NewGame() *Game {
	return &Game{
		ID:    uuid.NewString(),
		State: NewInitialState(),
	}
}

// ApplyMove mutates game state (NO validation here)

func (g *Game) ApplyMove(m Move) error {
	g.mu.Lock()
	defer g.mu.Unlock()

	if g.Ended {
		return errors.New("game already ended")
	}

	now := time.Now().UnixMilli()
	elapsed := now - g.State.LastTick

	if g.State.Turn == core.White {
		g.State.WhiteTime -= elapsed
	} else {
		g.State.BlackTime -= elapsed
	}

	g.State.LastTick = now

	ApplyMove(g.State, m)
	g.Moves = append(g.Moves, m)

	return nil
}

// ✅ ADD THIS METHOD
func (g *Game) Resign() {
	g.mu.Lock()
	defer g.mu.Unlock()
	g.Ended = true
}
