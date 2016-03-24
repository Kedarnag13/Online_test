package account

import (
	"database/sql"
	"encoding/json"
	"github.com/Kedarnag13/Online_test/api/v1/controllers"
	"github.com/Kedarnag13/Online_test/api/v1/models"
	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
	"io/ioutil"
	"log"
	"net/http"
)

type sessionController struct{}

var Session sessionController

func (s sessionController) Create(rw http.ResponseWriter, req *http.Request) {

	body, err := ioutil.ReadAll(req.Body)
	var u models.User

	if err != nil {
		panic(err)
	}
	err = json.Unmarshal(body, &u)
	if err != nil {
		panic(err)
	}
	
}
