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
if u.SectionId == 1 {
var insert_result string = "insert into results (user_id, section_1,) values ($1,$2)"

prepare_insert_result, err := db.Prepare(insert_result)
if err != nil {
	log.Fatal(err)
}

defer prepare_insert_result.Close()

insert_result_exec, err := prepare_insert_result.Exec(u.UserId,score)
if err != nil || insert_result_exec == nil {
	log.Fatal(err)
}

} else if u.SectionId == 2 {
	update_result, err := db.Query("UPDATE results SET section_2=$1 where user_id=$2", score, u.UserId)
	if err != nil || update_result == nil {
		log.Fatal(err)
	}
	defer update_result.Close()
	} else {
		fetch_score, err := db.Query("SELECT section_1, section_2 from results where user_id = $1")
		if err != nil || fetch_score == nil {
			log.Fatal(err)
		}
		defer fetch_score.Close()

		var section_1_score int
		var section_2_score int

		for fetch_score.Next() {
			err := fetch_score.Scan(&section_1_score, &section_2_score)
			if err != nil {
				log.Fatal(err)
			}
		}

		total_score := section_1_score + section_2_score + score

		update_result, err := db.Query("UPDATE results SET section_3=$1, total_score = $2 where user_id=$3", score, total_score, u.UserId)
		if err != nil || update_result == nil {
			log.Fatal(err)
		}
		defer update_result.Close()
		
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
