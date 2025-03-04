package gapi

import (
	"fmt"
	db "github.com/mitchxxx/simplebank/db/sqlc"
	"github.com/mitchxxx/simplebank/pb"
	"github.com/mitchxxx/simplebank/token"
	"github.com/mitchxxx/simplebank/util"
)

//Server serves HTTP request for our Banking service

type Server struct {
	pb.UnimplementedSimpleBankServer
	config util.Config
	store db.Store
	tokenMaker token.Maker
}

// NewServer creates a new GRPC server

func NewServer(config util.Config, store db.Store) (*Server, error) {
	tokenMaker, err := token.NewPasetoMaker(config.TokenSymmetricKey)
	if err != nil {
		return nil, fmt.Errorf("cannot create a token maker : %w", err)
	}
	server := &Server{
		config: config,
		store: store,
		tokenMaker: tokenMaker,
	}
	
	return server, nil
}
