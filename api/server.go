package api

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
	db "github.com/mitchxxx/simplebank/db/sqlc"
	"github.com/mitchxxx/simplebank/token"
	"github.com/mitchxxx/simplebank/util"
)

//Server serves HTTP request for our Banking service

type Server struct {
	config util.Config
	store db.Store
	router *gin.Engine
	tokenMaker token.Maker
}

// NewServer creates a new HTTP server and setup routing

func NewServer(config util.Config, store db.Store) (*Server, error) {
	tokenMaker, err := token.NewPasetoMaker(config.TokenSymmetricKey)
	//tokenMaker, err := token.NewJWTMaker(config.TokenSymmetricKey)
	if err != nil {
		return nil, fmt.Errorf("cannot create a token maker : %w", err)
	}
	server := &Server{
		config: config,
		store: store,
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

	// create a router group to add handlers that will be covered by the Authorization middleware
	authRoutes := router.Group("/").Use(authMiddleware(server.tokenMaker))
	authRoutes.POST("/accounts", server.createAccount)
	authRoutes.GET("/accounts/:id", server.getAccount)
	authRoutes.GET("/accounts", server.ListAccounts)
	authRoutes.POST("/transfers", server.CreateTransfer)
// Authorization middleware not required
	router.POST("/users", server.createUser)
	router.POST("/users/login", server.loginUser)
	router.POST("/tokens/renew_access", server.renewAccessToken)

	server.router = router
}

// Start runs the HTTP server on  a specific address
func (server *Server) Start(address string) error {
	return server.router.Run(address)
}

func errorResponse (err error) gin.H {
	return gin.H{"error": err.Error()}
}