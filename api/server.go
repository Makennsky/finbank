package api

import (
	db "github.com/Makennsky/finbank/db/sqlc"
	"github.com/gin-gonic/gin"
)

// Server serves Http requests for our service
type Server struct {
	store  *db.Store
	router *gin.Engine
}

// It creates new server instance
func NewServer(store *db.Store) *Server {
	server := &Server{store: store}
	router := gin.Default()

	router.POST("/accounts", server.createAccount)
	router.GET("/accounts", server.listAccount)
	router.GET("/accounts/:id", server.getAccount)

	server.router = router
	return server
}

// Start Http server
func (server *Server) Start(address string) error {
	return server.router.Run(address)
}

func errorResponse(err error) gin.H {
	return gin.H{"error": err.Error()}
}
