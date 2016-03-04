package exam

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"
	"time"
)

type examController struct{}

var Exam examController

func (e examController) Start(rw http.ResponseWriter, req *http.Request) {
	log.Println("Welcome")
}
