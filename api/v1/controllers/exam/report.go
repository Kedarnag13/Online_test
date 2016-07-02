package exam

import (
	"database/sql"
	"encoding/json"
	"github.com/Kedarnag13/Online_test/api/v1/models"
	_ "github.com/lib/pq"
	"net/http"
	"io/ioutil"
	"time"
	"log"
)

type resultController struct{}

var Result examController

func (e examController) Create(rw http.ResponseWriter, req *http.Request) {
	body, err := ioutil.ReadAll(req.Body)
	var u models.QuestionResponse

	flag := 0

	err = json.Unmarshal(body, &u)
	if err != nil {
		panic(err)
	}

	db, err := sql.Open("postgres", "password=password host=localhost dbname=online_test_dev sslmode=disable")
	if err != nil {
		panic(err)
	}

	get_section_id, err := db.Query("SELECT id from sections where name = $1", u.SectionName)
	if err != nil || get_section_id == nil {
		panic(err)
	}
	defer get_section_id.Close()

	var section_id int
	for get_section_id.Next(){
		err := get_section_id.Scan(&section_id)
		if err != nil {
			panic(err)
		}
	}

	score := 0
	log.Printf("result section_name: %v",u.SectionName)
	for _, i:= range u.Questions {
		log.Printf("result questions: %v",i.QuestionId)
		check_section, err := db.Query("SELECT answer FROM questions WHERE section_id = $1 AND id = $2",section_id,i.QuestionId)
		if err != nil || check_section == nil {
			panic(err)
		}
		defer check_section.Close()
		for check_section.Next(){
			log.Printf("inside check_section")
			var answer string
			err := check_section.Scan(&answer)
			if err != nil {
				panic(err)
			}
			if answer == i.Answer {
				log.Printf("correct answer")
				score = score + 1
			}
		}
	}

	user_details, err := db.Query("SELECT first_name, last_name, email, phone_number, city, batch from users where id = $1",u.UserId)
	if err != nil {
		panic(err)
	}
	defer user_details.Close()

	var first_name string
	var last_name string
	var email string
	var phone_number string
	var city string
	var batch string

	for user_details.Next(){
		log.Printf("fetching user details")
		err := user_details.Scan(&first_name, &last_name, &email, &phone_number, &city, &batch)
		if err != nil {
			panic(err)
		}
		check_result_exist, err := db.Query("SELECT user_id from results where user_id = $1", u.UserId)
		if err != nil {
			panic(err)
		}
		defer check_result_exist.Close()
		if section_id == 1 {
			start_time := time.Now()
			for check_result_exist.Next(){
				log.Printf("Updating Section 1 results")
				update_section1_results, err := db.Query("UPDATE results SET section_1= $1, start_time = $2 where user_id = $3", score, start_time, u.UserId)
				if err != nil || update_section1_results == nil {
					panic(err)
				}
				defer update_section1_results.Close()
				calculate_total_score(u.UserId)
				b, err := json.Marshal(models.Result{
					Section:  section_id,
					TotalQuestions: 20,
					Score:	score,
				})
				if err != nil {
					panic(err)
				}
				rw.Header().Set("Content-Type", "application/json")
				rw.Write(b)
				goto update_results_end
			}
			log.Printf("Creating and inserting Section 1 results")
			var insert_result string = "insert into results (user_id, section_1, first_name, last_name, email, phone_number, city, batch, start_time) values ($1,$2,$3,$4,$5,$6,$7,$8,$9)"
			prepare_insert_result, err := db.Prepare(insert_result)
			if err != nil {
				panic(err)
			}

			defer prepare_insert_result.Close()

			insert_result_exec, err := prepare_insert_result.Exec(u.UserId, score, first_name, last_name, email, phone_number, city, batch, start_time)
			if err != nil || insert_result_exec == nil {
				panic(err)
			}
			calculate_total_score(u.UserId)

			} else if section_id == 2 {
				start_time := time.Now()
				for check_result_exist.Next(){
					log.Printf("Updating Section 2 results")
					update_result, err := db.Query("UPDATE results SET section_2=$1 where user_id=$2", score, u.UserId)
					if err != nil || update_result == nil {
						panic(err)
					}
					defer update_result.Close()
					calculate_total_score(u.UserId)

					b, err := json.Marshal(models.Result{
						Section: section_id,
						TotalQuestions: 20,
						Score:	score,
					})
					if err != nil {
						panic(err)
					}
					rw.Header().Set("Content-Type", "application/json")
					rw.Write(b)
					goto update_results_end
				}
				log.Printf("Creating and inserting Section 2 results")
				var insert_result string = "insert into results (user_id, section_2, first_name, last_name, email, phone_number, city, batch, start_time) values ($1,$2,$3,$4,$5,$6,$7,$8,$9)"
				prepare_insert_result, err := db.Prepare(insert_result)
				if err != nil {
					panic(err)
				}

				defer prepare_insert_result.Close()

				insert_result_exec, err := prepare_insert_result.Exec(u.UserId, score, first_name, last_name, email, phone_number, city, batch, start_time)
				if err != nil || insert_result_exec == nil {
					panic(err)
				}
				calculate_total_score(u.UserId)
				} else {
					start_time := time.Now()
					for check_result_exist.Next(){
						log.Printf("Updating Section 3 results")
						update_result, err := db.Query("UPDATE results SET section_3=$1 where user_id=$2", score, u.UserId)
						if err != nil || update_result == nil {
							panic(err)
						}
						defer update_result.Close()
						calculate_total_score(u.UserId)

						b, err := json.Marshal(models.Result{
							Section: section_id,
							TotalQuestions: 20,
							Score:	score,
						})
						if err != nil {
							panic(err)
						}
						rw.Header().Set("Content-Type", "application/json")
						rw.Write(b)
						goto update_results_end
					}
					log.Printf("Creating and inserting Section 3 results")
					var insert_result string = "insert into results (user_id, section_3, first_name, last_name, email, phone_number, city, batch, start_time) values ($1,$2,$3,$4,$5,$6,$7,$8,$9)"
					prepare_insert_result, err := db.Prepare(insert_result)
					if err != nil {
						panic(err)
					}

					defer prepare_insert_result.Close()

					insert_result_exec, err := prepare_insert_result.Exec(u.UserId, score, first_name, last_name, email, phone_number, city, batch, start_time)
					if err != nil || insert_result_exec == nil {
						panic(err)
					}
					calculate_total_score(u.UserId)
				}
				b, err := json.Marshal(models.Result{
					Section: section_id,
					TotalQuestions: 20,
					Score:	score,
				})
				if err != nil {
					panic(err)
				}
				rw.Header().Set("Content-Type", "application/json")
				rw.Write(b)
				goto update_results_end
			}
			if flag == 0 {
				b, err := json.Marshal(models.ErrorMessage{
					Success: "false",
					Error:   "User does not exist",
				})

				if err != nil {
					panic(err)
				}
				rw.Header().Set("Content-Type", "application/json")
				rw.Write(b)
			}

			update_results_end:
			db.Close()
		}

		func calculate_total_score(user_id int){
			db, err := sql.Open("postgres", "password=password host=localhost dbname=online_test_dev sslmode=disable")
			if err != nil {
				panic(err)
			}

			log.Printf("Updating final results")
			fetch_score, err := db.Query("SELECT section_1, section_2, section_3 from results where user_id = $1", user_id)
			if err != nil || fetch_score == nil {
				panic(err)
			}
			defer fetch_score.Close()

			var section_1_score int
			var section_2_score int
			var section_3_score int

			for fetch_score.Next() {
				err := fetch_score.Scan(&section_1_score, &section_2_score, &section_3_score)
				if err != nil {
					panic(err)
				}
			}

			total_score := section_1_score + section_2_score + section_3_score
			end_time := time.Now()
			update_result, err := db.Query("UPDATE results SET total_score = $1, test_finished = $2, end_time = $3 where user_id = $4", total_score, true, end_time, user_id)
			if err != nil || update_result == nil {
				panic(err)
			}
			defer update_result.Close()
		}

		func (e examController) Export(rw http.ResponseWriter, req *http.Request) {
			db, err := sql.Open("postgres", "password=password host=localhost dbname=online_test_dev sslmode=disable")
			if err != nil {
				panic(err)
			}
			export_csv, err := db.Query("COPY results TO '/tmp/results.csv' DELIMITER ',' CSV HEADER;")
			if err != nil {
				panic(err)
			}
			defer export_csv.Close()
			db.Close()
		}

		func (e examController) ResultList(rw http.ResponseWriter, req *http.Request){
			db, err := sql.Open("postgres", "password=password host=localhost dbname=online_test_dev sslmode=disable")
			if err != nil {
				panic(err)
			}
			pass_list, err := db.Query("Select first_name, last_name, email, phone_number, city, batch, section_1, section_2, section_3, total_score, start_time, end_time from results where test_finished = true", )
			defer pass_list.Close()
			if err != nil {
				panic(err)
			}
			var all_user_results []models.UserResult
			var total_users int
			for pass_list.Next(){
				var first_name string
				var last_name string
				var email string
				var phone_number string
				var city string
				var batch string
				var section_1_score int
				var section_2_score int
				var section_3_score	int
				var total_score int
				var start_time time.Time
				var end_time time.Time
				var result models.UserResult
				err := pass_list.Scan(&first_name, &last_name, &email, &phone_number, &city, &batch, &section_1_score, &section_2_score, &section_3_score, &total_score, &start_time, &end_time)
				if err != nil {
					panic(err)
				}
				result = models.UserResult{first_name, last_name, email, phone_number, city, batch, section_1_score, section_2_score, section_3_score, total_score, start_time, end_time}
				all_user_results = append(all_user_results, result)
				total_users = total_users + 1
			}
			b, err := json.Marshal(models.Report{
				Total_users: total_users,
				Report_list:	all_user_results,
			})
			if err != nil {
				panic(err)
			}
			rw.Header().Set("Content-Type", "application/json")
			rw.Write(b)
			db.Close()
		}
