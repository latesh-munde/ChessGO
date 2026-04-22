package pieces

import "chess-server/internal/core"

func slidingMoves(
	from core.Square,
	board core.Board,
	color core.Color,
	dirs []struct{ r, c int },
) []core.Square {

	var moves []core.Square

	for _, d := range dirs {
		cur := from
		for {
			next, ok := cur.Offset(d.r, d.c)
			if !ok {
				break
			}
			if board.Empty(next) {
				moves = append(moves, next)
				cur = next
				continue
			}
			if board.EnemyAt(next, color) {
				moves = append(moves, next)
			}
			break
		}
	}
	return moves
}
