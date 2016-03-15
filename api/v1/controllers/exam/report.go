package exam

import (
	"database/sql"
	"encoding/json"
	"github.com/Kedarnag13/Online_test/api/v1/models"
	_ "github.com/lib/pq"
	"log"
	"net/http"
	"io/ioutil"
)

type resultController struct{}

var Result examController

func (e examController) Create(rw http.ResponseWriter, req *http.Request) {
	body, err := ioutil.ReadAll(req.Body)
	var u models.QuestionResponse
	
	err = json.Unmarshal(body, &u)
	if err != nil {
		panic(err)
	}

	db, err := sql.Open("postgres", "password=password host=localhost dbname=online_test_dev sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}

	score := 0

	for _, i:= range u.Questions {
		check_section, err := db.Query("SELECT answer FROM questions WHERE section_id = $1 AND id = $2",u.SectionId,i.QuestionId)
		if err != nil || check_section == nil {
			log.Fatal(err)
		}
		defer check_section.Close()
		for check_section.Next(){
			var answer string
			err := check_section.Scan(&answer)
			if err != nil {
				log.Fatal(err)
			}
			if answer == i.AnswerOption {
				score = score + 1
			}
		}
	}

	b, err := json.Marshal(models.Result{
		Section:     u.SectionId,
		TotalQuestions: 20,
		Score:	score,
	})
	if err != nil {
		log.Fatal(err)
	}
	rw.Header().Set("Content-Type", "application/json")
	rw.Write(b)

	db.Close()
}
