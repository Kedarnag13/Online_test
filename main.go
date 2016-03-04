package main

import (
	"github.com/Kedarnag13/Online_test/api/v1/controllers/account"
	"github.com/Kedarnag13/Online_test/api/v1/controllers/exam"
	"github.com/gorilla/mux"
	"github.com/rs/cors"
	"log"
	"net/http"
)

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/sign_up", account.Registration.Create).Methods("POST")
	r.HandleFunc("/start", exam.Exam.Start).Methods("GET")
	// HTTP Listening Port
	handler := cors.Default().Handler(r)
	http.Handle("/", handler)
	log.Println("main : Started : Listening on: http://localhost:3000 ...")
	log.Fatal(http.ListenAndServe("0.0.0.0:3000", nil))
}
