package main

import (
	"log"

	"github.com/gin-gonic/gin"

	"chess-server/internal/api"
	"chess-server/internal/store"
)

func main() {
	store := store.NewMemoryStore()

	r := gin.Default()
	api.RegisterRoutes(r, store)

	log.Println("Server running on http://localhost:8080")
	log.Fatal(r.Run(":8080"))
}
