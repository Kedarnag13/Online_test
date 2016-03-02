package main

import (
	"github.com/gorilla/mux"
	"github.com/kedarnag13/Online_test/app/v1/controllers/account"
	"log"
	"net/http"
)

func main() {

	r := mux.NewRouter()
	r.HandleFunc("/sign_up", account.User.Create).Methods("POST")
	http.Handle("/", r)

	// HTTP Listening Port
	log.Println("main : Started : Listening on: http://localhost:3000 ...")
	log.Fatal(http.ListenAndServe("0.0.0.0:3000", nil))

}
