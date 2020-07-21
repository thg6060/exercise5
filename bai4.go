package main

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/thg6060/exercise5/database"
	"github.com/thg6060/exercise5/protoc"
)

var api = database.UserPartApi{}

func GetList(w http.ResponseWriter, req *http.Request) {

	w.Header().Set("Content-Type", "application/json")
	data, err := api.GetUserList()
	if err != nil {
		panic(err)
	}
	json.NewEncoder(w).Encode(data)
}
func DeleteUsr(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	message := "Success"
	params := mux.Vars(req)
	err := api.DeleteUser(params["id"])
	if err != nil {
		panic(err)
	}
	json.NewEncoder(w).Encode(message)
}

func GetUsrById(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(req)
	data, err := api.GetUserById(params["id"])
	if err != nil {
		panic(err)
	}
	json.NewEncoder(w).Encode(data)
}
func CreateUsr(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	message := "Success"
	usr := protoc.UserPartner{}
	data, _ := ioutil.ReadAll(req.Body)
	err := json.Unmarshal(data, &usr)
	if err != nil {
		panic(err)
	}
	err = api.CreateUser(&usr)

	if err != nil {
		panic(err)
	}

	json.NewEncoder(w).Encode(message)
}

func InitRoute() *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/user-partner", GetList).Methods("GET")
	r.HandleFunc("/user-partner/{id}", DeleteUsr).Methods("DELETE")
	r.HandleFunc("/user-partner/{id}", GetUsrById).Methods("GET")
	r.HandleFunc("/user-partner", CreateUsr).Methods("POST")
	return r
}

func InitServer() error {

	r := InitRoute()
	err := http.ListenAndServe(":3011", r)
	if err != nil {
		return err
	}
	return nil
}
