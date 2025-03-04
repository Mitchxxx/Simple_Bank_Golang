package main

import (
	"database/sql"
	"log"
	"net"

	_ "github.com/lib/pq"
	"github.com/mitchxxx/simplebank/api"
	db "github.com/mitchxxx/simplebank/db/sqlc"
	"github.com/mitchxxx/simplebank/gapi"
	"github.com/mitchxxx/simplebank/pb"
	"github.com/mitchxxx/simplebank/util"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)


func main(){
	config, err := util.LoadConfig(".")
	if err != nil {
		log.Fatal("cannot load config:", err)
	}
	conn, err := sql.Open(config.DBDriver, config.DBSource)
	if err != nil {
		log.Fatal("cannot connect to db:", err)
	}

	store := db.NewStore(conn)
	//runGinServer(config, store)
	runGrpcServer(config, store)

}

func runGrpcServer(config util.Config, store db.Store) {
	server, err := gapi.NewServer(config, store)
	if err != nil {
		log.Fatal("Cannot create server:", err)
	}
	grpcServer := grpc.NewServer()
	pb.RegisterSimpleBankServer(grpcServer, server)
	reflection.Register(grpcServer)

	listener, err := net.Listen("tcp", config.GRPCServerAddress)
	if err != nil {
		log.Fatal("cannot create listerner:", err)
	}
	log.Printf("Start gRPC server on %s", listener.Addr().String())
	err = grpcServer.Serve(listener)
	if err!= nil {
		log.Fatal("Cannot start gRPC server:", err)
	}


}

func runGinServer(config util.Config, store db.Store) {

	server, err := api.NewServer(config,store)
	if err != nil {
        log.Fatal("cannot create server:", err)
    }
	err = server.Start(config.HTTPServerAddress)
	if err != nil {
		log.Fatal("cannot start server:", err)
	}
}