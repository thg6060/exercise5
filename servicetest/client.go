package test

import (
	"context"
	"fmt"
	"testing"

	"github.com/thg6060/exercise5/protoc"
	"google.golang.org/grpc"
)

func TestRunClient(t *testing.T) {
	RunClient()
}

func RunClient() error {

	conn, err := grpc.Dial(":7777", grpc.WithInsecure())
	defer conn.Close()
	if err != nil {
		return err
	}

	service := protoc.NewUserPartnerServiceClient(conn)

	req := &protoc.UserPartnerListRequest{}

	data, err := service.GetUserPartnerList(context.Background(), req)

	for _, item := range data.ListUser {
		fmt.Println(item)
	}

	/*
		data, err := service.GetUserPartnerById(context.Background(), req)
		fmt.Println(data)

	*/

	if err != nil {
		return err
	}
	return nil
}
