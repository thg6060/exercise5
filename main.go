package main
/*
package main

import (
	//"fmt"

	//"github.com/rs/xid"

	"fmt"
	"io"
	"log"
	"net/http"

	"github.com/thg6060/exercise5/database"
	"github.com/thg6060/exercise5/grpc"
	"github.com/thg6060/exercise5/protoc"
)

var db = database.DbConn()

func Bai2() {
	obj := database.UserPartApi{}

	condition := protoc.UserPartnerByIdRequest{
		UserId: "bsali7rqgn6oggb2h1t0",
	}
	data, err := obj.GetUserByCdt(&condition)

	if err != nil {
		fmt.Println(err)
	}
	for _, item := range data {
		fmt.Println(item)
	}
}
func Bai3() {

	helloHandler := func(w http.ResponseWriter, req *http.Request) {
		io.WriteString(w, "pong!\n")
	}
	http.HandleFunc("/", helloHandler)
	log.Fatal(http.ListenAndServe(":3003", nil))

}
/*
func main() {
	//database.CreateTable()
	//Bai3()
	var api protoc.UserPartnerServiceServer
	grpc.RunServer(api)

}
*/
