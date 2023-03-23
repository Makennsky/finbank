package api

import (
	"fmt"

	db "github.com/Makennsky/finbank/db/sqlc"
	"github.com/Makennsky/finbank/token"
	"github.com/Makennsky/finbank/util"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
)

// Server serves Http requests for service
type Server struct {
	config     util.Config
	store      db.Store
	tokenMaker token.Maker
	router     *gin.Engine
}

// It creates new server instance
func NewServer(conf util.Config, store db.Store) (*Server, error) {
	tokenMaker, err := token.NewPasetoMaker(conf.TokenSymmetricKey)
	if err != nil {
		return nil, fmt.Errorf("cannot create a token maker: %w", err)
	}
	server := &Server{
		config:     conf,
		store:      store,
		tokenMaker: tokenMaker,
	}
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		v.RegisterValidation("currency", validCurrency)
	}

	server.setupRouter()
	return server, nil
}

func (server *Server) setupRouter() {
	router := gin.Default()

	router.POST("/users", server.createUser)
	router.GET("/users/:username", server.getUser)
	router.POST("/users/login", server.loginUser)

	authRoutes := router.Group("/").Use(authMiddleware(server.tokenMaker))

	authRoutes.POST("/accounts", server.createAccount)
	authRoutes.GET("/accounts", server.listAccount)
	authRoutes.GET("/accounts/:id", server.getAccount)

	authRoutes.POST("/transfers", server.createTransfer)

	server.router = router
}

// Start Http server
func (server *Server) Start(address string) error {
	return server.router.Run(address)
}

func errorResponse(err error) gin.H {
	return gin.H{"error": err.Error()}
}
