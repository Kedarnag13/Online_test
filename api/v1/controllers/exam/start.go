package exam

import (
	"database/sql"
	"encoding/json"
	"github.com/Kedarnag13/Online_test/api/v1/models"
	_ "github.com/lib/pq"
	"log"
	"net/http"
)

type examController struct{}

var Exam examController

func (e examController) Start(rw http.ResponseWriter, req *http.Request) {
	db, err := sql.Open("postgres", "password=password host=localhost dbname=online_test_dev sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}

	dumps, err := db.Exec("CREATE TABLE IF NOT EXISTS dumps (data json)")
	if err != nil || dumps == nil {
		log.Fatal(err)
	}

	questions, err := db.Query("SELECT data FROM dumps")
	if err != nil || questions == nil {
		log.Fatal(err)
	}
	defer questions.Close()
	for questions.Next() {
		var data string
		err = questions.Scan(&data)
		if err != nil {
			log.Println(err)
		}
		b, err := json.Marshal(models.Data{
			Data: data,
		})
		if err != nil {
			log.Fatal(err)
		}
		rw.Header().Set("Content-Type", "application/json")
		rw.Write(b)
	}
}
