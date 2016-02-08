package main

import (
	"github.com/gorilla/mux"
	"github.com/kedarnag13/Online_test/api/v1/controllers/account"
	"log"
	"net/http"
)

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/sign_up", account.User.Create).Methods("POST")
	// HTTP Listening Port
	log.Println("main : Started : Listening on: http://localhost:3020 ...")
	log.Fatal(http.ListenAndServe("0.0.0.0:3020", nil))
}
