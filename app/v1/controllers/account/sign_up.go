package account

import (
	"log"
	"net/http"
)

type registrationController struct{}

var User registrationController

func (r registrationController) Create(rw http.ResponseWriter, req *http.Request) {
	log.Println("Welcome to Online Test")
}
