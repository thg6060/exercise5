package main

import (
	"net"

	"github.com/thg6060/exercise5/protoc"
	"google.golang.org/grpc"
)

func RunServer(api protoc.UserPartnerServiceServer) error {

	listen, err := net.Listen("tcp", ":7777")
	if err != nil {
		return err
	}
	server := grpc.NewServer()
	var us protoc.UserPartnerSerice
	protoc.RegisterUserPartnerServiceServer(server, &us)

	server.Serve(listen)
	return nil
}
func main() {
	var api protoc.UserPartnerServiceServer
	RunServer(api)
}
