package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"

	"github.com/thg6060/exercise5/database"
	"github.com/thg6060/exercise5/protoc"
	"github.com/urfave/cli"
)



func Bai2() {

	obj := database.UserPartApi{}

	condition := protoc.UserPartnerByIdRequest{
		UserId: "bsali7rqgn6oggb2h1t0",
		Limit: 5,
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

func Menu(order string) {

	switch order {
	case "bai2":
		Bai2()
	case "bai3":
		Bai3()
	}
}

func MenuCli(c *cli.Context) {
	Menu(c.Args().Get(0))
}

func main() {

	app := &cli.App{
		Name:   "bing",
		Usage:  "grpc",
		Action: MenuCli,
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}

}
