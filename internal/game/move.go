package game

import (
	"errors"

	"chess-server/internal/core"
)

type Move struct {
	From core.Square
	To   core.Square
	Prom string
}

func ParseMove(from, to, prom string) (Move, error) {
	if len(from) != 2 || len(to) != 2 {
		return Move{}, errors.New("invalid square format")
	}

	fFile := int(from[0] - 'a')
	fRank := int(from[1] - '1')
	tFile := int(to[0] - 'a')
	tRank := int(to[1] - '1')

	if fFile < 0 || fFile > 7 || fRank < 0 || fRank > 7 ||
		tFile < 0 || tFile > 7 || tRank < 0 || tRank > 7 {
		return Move{}, errors.New("square out of bounds")
	}

	return Move{
		From: core.Square{Rank: fRank, File: fFile},
		To:   core.Square{Rank: tRank, File: tFile},
		Prom: prom,
	}, nil
}
