package game

import (
	"chess-server/internal/core"
	"chess-server/internal/pieces"
	"time"
)

type State struct {
	Board *Board
	Turn  core.Color

	// ⏱ CLOCKS (milliseconds)
	WhiteTime int64
	BlackTime int64

	LastTick  int64 // unix millis
	EnPassant *core.Square
}

type CastlingRights struct {
	WhiteKing, WhiteQueen bool
	BlackKing, BlackQueen bool
}

func NewInitialState() *State {
	b := NewBoard()

	// Pawns
	for f := 0; f < 8; f++ {
		b.Squares[1][f] = pieces.NewPawn(core.White)
		b.Squares[6][f] = pieces.NewPawn(core.Black)
	}

	// Rooks
	b.Squares[0][0] = pieces.NewRook(core.White)
	b.Squares[0][7] = pieces.NewRook(core.White)
	b.Squares[7][0] = pieces.NewRook(core.Black)
	b.Squares[7][7] = pieces.NewRook(core.Black)

	// Knights
	b.Squares[0][1] = pieces.NewKnight(core.White)
	b.Squares[0][6] = pieces.NewKnight(core.White)
	b.Squares[7][1] = pieces.NewKnight(core.Black)
	b.Squares[7][6] = pieces.NewKnight(core.Black)

	// Bishops
	b.Squares[0][2] = pieces.NewBishop(core.White)
	b.Squares[0][5] = pieces.NewBishop(core.White)
	b.Squares[7][2] = pieces.NewBishop(core.Black)
	b.Squares[7][5] = pieces.NewBishop(core.Black)

	// Queens
	b.Squares[0][3] = pieces.NewQueen(core.White)
	b.Squares[7][3] = pieces.NewQueen(core.Black)

	// Kings
	b.Squares[0][4] = pieces.NewKing(core.White)
	b.Squares[7][4] = pieces.NewKing(core.Black)

	now := time.Now().UnixMilli()

	return &State{
		Board:     b,
		Turn:      core.White,
		WhiteTime: 5 * 60 * 1000, // 5 minutes
		BlackTime: 5 * 60 * 1000,
		LastTick:  now,
		EnPassant: nil,
	}
}
