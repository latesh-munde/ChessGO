package core

type Board interface {
	Empty(Square) bool
	EnemyAt(Square, Color) bool
	FriendlyAt(Square, Color) bool
}
