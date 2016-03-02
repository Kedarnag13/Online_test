package main

import (
	"github.com/Kedarnag13/Online_test/api/v1/controllers/account"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/sign_up", account.User.Create).Methods("POST")
	// HTTP Listening Port

	http.Handle("/", r)
	log.Println("main : Started : Listening on: http://localhost:3000 ...")
	log.Fatal(http.ListenAndServe("0.0.0.0:3000", nil))
}
