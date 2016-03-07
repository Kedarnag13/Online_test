package exam

import (
	"database/sql"
	"encoding/json"
	"github.com/Kedarnag13/Online_test/api/v1/models"
	_ "github.com/lib/pq"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"strconv"
)

type examController struct{}

var Exam examController

func (e examController) Logical(rw http.ResponseWriter, req *http.Request) {

	var l models.Aptitude
	vars := mux.Vars(req)
	id := vars["id"]
	tmp, err := strconv.Atoi(id)
	l.Id = tmp

	db, err := sql.Open("postgres", "password=password host=localhost dbname=online_test_dev sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}

	logicals, err := db.Exec("CREATE TABLE IF NOT EXISTS logicals (id SERIAL, data jsonb)")
	if err != nil || logicals == nil {
		log.Fatal(err)
	}

	questions, err := db.Query("SELECT data FROM logicals WHERE id=$1", l.Id)
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

	var a models.Aptitude
	vars := mux.Vars(req)
	id := vars["id"]
	tmp, err := strconv.Atoi(id)
	a.Id = tmp

	db, err := sql.Open("postgres", "password=password host=localhost dbname=online_test_dev sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}

	aptitudes, err := db.Exec("CREATE TABLE IF NOT EXISTS aptitudes (id SERIAL, data jsonb)")
	if err != nil || aptitudes == nil {
		log.Fatal(err)
	}

	questions, err := db.Query("SELECT data FROM aptitudes WHERE id=$1", a.Id)
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

	var v models.Verbal
	vars := mux.Vars(req)
	id := vars["id"]
	tmp, err := strconv.Atoi(id)
	v.Id = tmp

	db, err := sql.Open("postgres", "password=password host=localhost dbname=online_test_dev sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}

	verbals, err := db.Exec("CREATE TABLE IF NOT EXISTS verbals (id SERIAL, data jsonb)")
	if err != nil || verbals == nil {
		log.Fatal(err)
	}

	questions, err := db.Query("SELECT data FROM verbals WHERE id=$1", v.Id)
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
