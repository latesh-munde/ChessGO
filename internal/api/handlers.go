package api

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"chess-server/internal/engine"
	"chess-server/internal/game"
	"chess-server/internal/store"
)

type Handler struct {
	store store.GameStore
}

func NewHandler(s store.GameStore) *Handler {
	return &Handler{store: s}
}

func (h *Handler) CreateGame(c *gin.Context) {
	g := game.NewGame()
	h.store.Save(g)

	c.JSON(http.StatusCreated, gin.H{
		"id": g.ID,
	})
}

func (h *Handler) JoinGame(c *gin.Context) {
	id := c.Param("id")

	if err := h.store.Join(id); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.Status(http.StatusOK)
}

func (h *Handler) MakeMove(c *gin.Context) {
	id := c.Param("id")

	var req struct {
		From      string `json:"from"`
		To        string `json:"to"`
		Promotion string `json:"promotion"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid request"})
		return
	}

	g, err := h.store.Get(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "game not found"})
		return
	}

	move, err := game.ParseMove(req.From, req.To, req.Promotion)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// ✅ VALIDATION LIVES IN ENGINE
	if err := engine.ValidateMove(g.State, move); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// ✅ GAME ONLY APPLIES MOVE
	if err := g.ApplyMove(move); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.Status(http.StatusOK)
}

func (h *Handler) GetState(c *gin.Context) {
	id := c.Param("id")

	g, err := h.store.Get(id)
	if err != nil {
		c.JSON(404, gin.H{"error": "game not found"})
		return
	}

	var board [8][8]gin.H
	for r := 0; r < 8; r++ {
		for f := 0; f < 8; f++ {
			p := g.State.Board.Squares[r][f]
			if p != nil {
				board[r][f] = gin.H{
					"type":  p.Type().String(),
					"color": p.Color(),
				}
			}
		}
	}

	var lastMove gin.H
	if len(g.Moves) > 0 {
		m := g.Moves[len(g.Moves)-1]
		lastMove = gin.H{
			"from": m.From.String(),
			"to":   m.To.String(),
		}
	}

	c.JSON(200, gin.H{
		"turn":      g.State.Turn,
		"board":     board,
		"inCheck":   engine.InCheck(g.State, g.State.Turn),
		"checkmate": engine.IsCheckmate(g.State, g.State.Turn),
		"stalemate": engine.IsStalemate(g.State, g.State.Turn),
		// ⏱ clocks
		"whiteTime": g.State.WhiteTime,
		"blackTime": g.State.BlackTime,
		"lastMove":  lastMove,
	})
}

func (h *Handler) GetMoves(c *gin.Context) {
	id := c.Param("id")

	g, err := h.store.Get(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "game not found"})
		return
	}

	c.JSON(http.StatusOK, g.Moves)
}

func (h *Handler) Resign(c *gin.Context) {
	id := c.Param("id")

	g, err := h.store.Get(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "game not found"})
		return
	}

	g.Resign()
	c.Status(http.StatusOK)
}

func (h *Handler) GetLegalMoves(c *gin.Context) {
	id := c.Param("id")
	from := c.Query("from")

	if from == "" {
		c.JSON(400, gin.H{"error": "from query param required"})
		return
	}

	g, err := h.store.Get(id)
	if err != nil {
		c.JSON(404, gin.H{"error": "game not found"})
		return
	}

	move, err := game.ParseMove(from, from, "")
	if err != nil {
		c.JSON(400, gin.H{"error": "invalid square"})
		return
	}

	p := g.State.Board.Squares[move.From.Rank][move.From.File]
	if p == nil {
		c.JSON(200, []string{})
		return
	}

	legal := p.LegalMoves(move.From, g.State.Board)

	resp := make([]string, 0, len(legal))
	for _, sq := range legal {
		resp = append(resp, sq.String())
	}

	c.JSON(200, resp)
}
