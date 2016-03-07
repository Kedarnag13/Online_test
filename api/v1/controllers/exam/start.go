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

	db, err := sql.Open("postgres", "password=password host=localhost dbname=online_test_dev sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}

	verbals, err := db.Exec("CREATE TABLE IF NOT EXISTS verbals (data jsonb)")
	if err != nil || verbals == nil {
		log.Fatal(err)
	}

	questions, err := db.Query("SELECT data FROM verbals")
	if err != nil || questions == nil {
		log.Fatal(err)
	}
	defer questions.Close()

	columns, err := questions.Columns() // Will return the columns
  if err != nil {
      log.Fatal(err)
  }
  count := len(columns) // Length of the columns
  tableData := make([]map[string]interface{}, 0) // keys to those columns
  values := make([]interface{}, count) // values of the respective column
  valuePtrs := make([]interface{}, count) // pointing to those columns
  for questions.Next() {
      for i := 0; i < count; i++ {
          valuePtrs[i] = &values[i]
      }
      questions.Scan(valuePtrs...)
      entry := make(map[string]interface{})
      for i, col := range columns {
          var v interface{}
          val := values[i]
          b, ok := val.([]byte)
          if ok {
              v = string(b)
          } else {
              v = val
          }
          entry[col] = v
      }
      tableData = append(tableData, entry) // Adding every entry that we get from the table into tableData
  }
  jsonData, err := json.Marshal(tableData)
  if err != nil {
      log.Fatal(err)
  }
	b, err := json.Marshal(models.Verbal{
			Data: string(jsonData),
		})
		if err != nil {
			log.Fatal(err)
		}
		rw.Header().Set("Content-Type", "application/json")
		rw.Write(b)
}
