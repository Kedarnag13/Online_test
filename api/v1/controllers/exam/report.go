package exam

import (
	"database/sql"
	"encoding/json"
	"github.com/Kedarnag13/Online_test/api/v1/models"
	_ "github.com/lib/pq"
	"net/http"
	"io/ioutil"
	"fmt"
	"log"
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
		panic(err)
	}

	score := 0
	fmt.Println("section_id:",u.SectionId)
	for _, i:= range u.Questions {
		fmt.Println("questions:",i.QuestionId)
		check_section, err := db.Query("SELECT answer FROM questions WHERE section_id = $1 AND id = $2",u.SectionId,i.QuestionId)
		if err != nil || check_section == nil {
			panic(err)
		}
		defer check_section.Close()
		for check_section.Next(){
			var answer string
			err := check_section.Scan(&answer)
			if err != nil {
				panic(err)
			}
			if answer == i.Answer {
				score = score + 1
			}
		}
	}

	user_details, err := db.Query("SELECT first_name, last_name, email from users where id = $1",u.UserId)
	if err != nil {
		panic(err)
	}
	defer user_details.Close()
	var first_name string
	var last_name string
	var email string

	for user_details.Next(){
		err := user_details.Scan(&first_name, &last_name, &email)
		if err != nil {
			panic(err)
		}
	}
	if u.SectionId == 1 {
		var insert_result string = "insert into results (user_id, section_1, first_name, last_name, email) values ($1,$2,$3,$4,$5)"

		prepare_insert_result, err := db.Prepare(insert_result)
		if err != nil {
			panic(err)
		}

		defer prepare_insert_result.Close()

		insert_result_exec, err := prepare_insert_result.Exec(u.UserId,score, first_name, last_name, email)
		if err != nil || insert_result_exec == nil {
			panic(err)
		}

		} else if u.SectionId == 2 {
			update_result, err := db.Query("UPDATE results SET section_2=$1 where user_id=$2", score, u.UserId)
			if err != nil || update_result == nil {
				panic(err)
			}
			defer update_result.Close()
			} else {
				fetch_score, err := db.Query("SELECT section_1, section_2 from results where user_id = $1", u.UserId)
				if err != nil || fetch_score == nil {
					panic(err)
				}
				defer fetch_score.Close()

				var section_1_score int
				var section_2_score int

				for fetch_score.Next() {
					err := fetch_score.Scan(&section_1_score, &section_2_score)
					if err != nil {
						panic(err)
					}
				}

				total_score := section_1_score + section_2_score + score

				update_result, err := db.Query("UPDATE results SET section_3=$1, total_score = $2 where user_id=$3", score, total_score, u.UserId)
				if err != nil || update_result == nil {
					panic(err)
				}
				defer update_result.Close()

			}

			b, err := json.Marshal(models.Result{
				Section:     u.SectionId,
				TotalQuestions: 20,
				Score:	score,
			})
			if err != nil {
				panic(err)
			}
			rw.Header().Set("Content-Type", "application/json")
			rw.Write(b)

			db.Close()
		}

func (e examController) Export(rw http.ResponseWriter, req *http.Request) {
	db, err := sql.Open("postgres", "password=password host=localhost dbname=online_test_dev sslmode=disable")
	if err != nil {
		panic(err)
	}
	export_csv, err := db.Query("COPY results TO '/Users/kedarnag/results.csv' DELIMITER ',' CSV HEADER;")
	if err != nil {
		panic(err)
	}
	defer export_csv.Close()
	db.Close()
}
