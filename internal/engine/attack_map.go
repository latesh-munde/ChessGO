package engine

import (
	"chess-server/internal/core"
	"chess-server/internal/game"
)

// AttackMap returns all squares attacked by a given color
func AttackMap(state *game.State, color core.Color) map[core.Square]bool {
	attacked := make(map[core.Square]bool)

	for r := 0; r < 8; r++ {
		for f := 0; f < 8; f++ {
			p := state.Board.Squares[r][f]
			if p == nil || p.Color() != color {
				continue
			}

			from := core.Square{Rank: r, File: f}
			moves := p.LegalMoves(from, state.Board)

			for _, sq := range moves {
				attacked[sq] = true
			}
		}
	}

	return attacked
}
