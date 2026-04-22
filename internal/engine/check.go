package engine

import (
	"chess-server/internal/core"
	"chess-server/internal/game"
	"chess-server/internal/pieces"
)

// InCheck returns true if the given color's king is under attack
func InCheck(state *game.State, color core.Color) bool {
	var kingSquare *core.Square

	// Find king position
	for r := 0; r < 8; r++ {
		for f := 0; f < 8; f++ {
			p := state.Board.Squares[r][f]
			if p == nil {
				continue
			}

			if p.Type() == pieces.King && p.Color() == color {
				ks := core.Square{Rank: r, File: f}
				kingSquare = &ks
				break
			}
		}
	}

	if kingSquare == nil {
		return false
	}

	// Get all enemy attack squares
	enemyAttacks := AttackMap(state, color.Opponent())
	return enemyAttacks[*kingSquare]
}
