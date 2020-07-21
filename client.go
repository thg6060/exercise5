package main

import (
	"context"
	"fmt"

	"github.com/thg6060/exercise5/protoc"
	"google.golang.org/grpc"
)

func RunClient() error {

	conn, err := grpc.Dial(":7777", grpc.WithInsecure())
	defer conn.Close()
	if err != nil {
		return err
	}

	service := protoc.NewUserPartnerServiceClient(conn)

	//req := &protoc.UserPartnerListRequest{}

     req:=&protoc.DeleteUserPartneRequest{
		 Id:"bsali7rqgn6oggb2h4324",
	 }
	
	/*
		data, err := service.GetUserPartnerList(context.Background(), req)


		for _, item := range data.ListUser {
			fmt.Println(item)
		}

	*/

	/*
	data, err := service.GetUserPartnerById(context.Background(), req)
	fmt.Println(data)

	*/
	data, err := service.DeleteUserPartnerById(context.Background(), req)
	fmt.Println(data.Status)

	if err != nil {
		return err
	}
	return nil
}
func main() {
	RunClient()
}
