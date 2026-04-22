package engine

import "chess-server/internal/game"

// IsDraw is a placeholder for full draw rules
func IsDraw(state *game.State) bool {
	// TODO:
	// - insufficient material
	// - repetition
	// - 50 move rule
	return false
}
