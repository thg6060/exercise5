package test

import (
	"net"
	"testing"

	"github.com/thg6060/exercise5/protoc"
	"google.golang.org/grpc"
)

func TestRunServer(t *testing.T){
	RunServer()
}

func RunServer() error {

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

