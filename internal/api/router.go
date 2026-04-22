package api

import (
	"chess-server/internal/store"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(r *gin.Engine, store store.GameStore) {
	r.Use(RequestLogger())

	RegisterUI(r) // 🔴 MUST BE HERE

	h := NewHandler(store)

	g := r.Group("/games")
	{
		g.POST("", h.CreateGame)
		g.POST("/:id/join", h.JoinGame)
		g.POST("/:id/move", h.MakeMove)
		g.GET("/:id/state", h.GetState)
		g.GET("/:id/moves", h.GetMoves)
		g.GET("/:id/legal", h.GetLegalMoves)
		g.POST("/:id/resign", h.Resign)

	}
}
