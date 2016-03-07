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

func (e examController) Logical(rw http.ResponseWriter, req *http.Request) {
	db, err := sql.Open("postgres", "password=password host=localhost dbname=online_test_dev sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}

	dumps, err := db.Exec("CREATE TABLE IF NOT EXISTS logicals (data jsonb)")
	if err != nil || dumps == nil {
		log.Fatal(err)
	}

	questions, err := db.Query("SELECT data FROM logicals")
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
		b, err := json.Marshal(models.Logical{
			Data: data,
		})
		if err != nil {
			log.Fatal(err)
		}
		rw.Header().Set("Content-Type", "application/json")
		rw.Write(b)
	}
}

func (e examController) Aptitude(rw http.ResponseWriter, req *http.Request) {
	db, err := sql.Open("postgres", "password=password host=localhost dbname=online_test_dev sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}

	dumps, err := db.Exec("CREATE TABLE IF NOT EXISTS aptitudes (data jsonb)")
	if err != nil || dumps == nil {
		log.Fatal(err)
	}

	questions, err := db.Query("SELECT data FROM aptitudes")
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
		b, err := json.Marshal(models.Aptitude{
			Data: data,
		})
		if err != nil {
			log.Fatal(err)
		}
		rw.Header().Set("Content-Type", "application/json")
		rw.Write(b)
	}
}

func (e examController) Verbal(rw http.ResponseWriter, req *http.Request) {
	db, err := sql.Open("postgres", "password=password host=localhost dbname=online_test_dev sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}

	dumps, err := db.Exec("CREATE TABLE IF NOT EXISTS verbals (data jsonb)")
	if err != nil || dumps == nil {
		log.Fatal(err)
	}

	questions, err := db.Query("SELECT data FROM verbals")
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
		b, err := json.Marshal(models.Verbal{
			Data: data,
		})
		if err != nil {
			log.Fatal(err)
		}
		rw.Header().Set("Content-Type", "application/json")
		rw.Write(b)
	}
}
