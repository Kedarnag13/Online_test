package exam

import (
"database/sql"
"encoding/json"
"github.com/Kedarnag13/Online_test/api/v1/models"
_ "github.com/lib/pq"
"net/http"
"io/ioutil"
"log"
"github.com/gorilla/mux"
)

type questionController struct{}

var Question questionController

func (q questionController) Create(rw http.ResponseWriter, req *http.Request) {
  body, err := ioutil.ReadAll(req.Body)
  if err != nil {
    panic(err)
  }

  var u models.CreateQuestion

  err = json.Unmarshal(body, &u)
  if err != nil {
    panic(err)
  }

  db, err := sql.Open("postgres", "password=password host=localhost dbname=online_test_dev sslmode=disable")
  if err != nil {
    panic(err)
  }
  questions, err:= db.Exec("CREATE TABLE IF NOT EXISTS questions(id int, title text, option_1 varchar(100), option_2 varchar(100), option_3 varchar(100), option_4 varchar(100), answer varchar(100), section_id int, CONSTRAINT section_id_key FOREIGN KEY(section_id) REFERENCES sections (id), PRIMARY KEY(id))")
  if err != nil || questions == nil {
    panic(err)
  }

  get_question_id, err := db.Query("SELECT coalesce(max(id), 0) FROM questions")
  if err != nil || get_question_id == nil {
    panic(err)
  }

  var section_id int
  switch u.Section {
    case "Verbal": section_id = 1
    case "Logical": section_id = 2
    case "Aptitude": section_id = 3

  }

  log.Printf("section_id: %v",section_id)
  var question_id int
  for get_question_id.Next(){
    err = get_question_id.Scan(&question_id)
    if err !=nil {
      panic(err)  
    }
    question_id = question_id + 1 

    var insert_question string = "insert into questions(id, title, option_1, option_2, option_3, option_4, answer, section_id) values($1,$2,$3,$4,$5,$6,$7,$8)"

    stmt, err := db.Prepare(insert_question)
    if err != nil {
      panic(err)
    }
    defer stmt.Close()

    insert_questions_exec, err := stmt.Exec(question_id, u.Question, u.OptionA, u.OptionB, u.OptionC, u.OptionD, u.Answer, section_id)
    if err != nil || insert_questions_exec == nil {
      panic(err)
    }
    log.Printf("question inserted successfully")
    b, err := json.Marshal(models.CreateQuestionStatusMessage{
      Success: "true",
      Message: "Question inserted Successfully!",
      })
    if err != nil {
      panic(err)
    }
    rw.Header().Set("Content-Type", "application/json")
    rw.Write(b)
  }
}

func (q questionController) AllQuestions(rw http.ResponseWriter, req *http.Request) {


  db, err := sql.Open("postgres", "password=password host=localhost dbname=online_test_dev sslmode=disable")
  if err != nil {
    panic(err)
  }
  questions, err:= db.Exec("CREATE TABLE IF NOT EXISTS questions(id int, title text, option_1 varchar(100), option_2 varchar(100), option_3 varchar(100), option_4 varchar(100), answer varchar(100), section_id int, CONSTRAINT section_id_key FOREIGN KEY(section_id) REFERENCES sections (id), PRIMARY KEY(id))")
  if err != nil || questions == nil {
    panic(err)
  }

  get_all_questions, err := db.Query("SELECT * from questions")
  if err != nil || get_all_questions == nil {
    panic(err)
  }
  all_questions := []models.FetchQuestion{}

  for get_all_questions.Next(){
    var Id int
    var Question string
    var Option_a string
    var Option_b string
    var Option_c string
    var Option_d string
    var Answer string
    var Section_id int

    var each_question models.FetchQuestion

    err := get_all_questions.Scan(&Id, &Question, &Option_a, &Option_b, &Option_c, &Option_d, &Answer, &Section_id)
    if err != nil {
      panic(err)
    }

    each_question = models.FetchQuestion{Id, Question, Option_a, Option_b, Option_c, Option_d, Answer, Section_id}
    all_questions = append(all_questions, each_question)
  }
  b, err := json.Marshal(models.FetchQuestionResponseMessage{
    Success:     "true",
    Message:     "Questions per section",
    QuestionList: all_questions,
    })
  if err != nil {
    panic(err)
  }
  rw.Header().Set("Content-Type", "application/json")
  rw.Write(b)
  db.Close()
}

func (q questionController) Edit(rw http.ResponseWriter, req *http.Request) {

  body, err := ioutil.ReadAll(req.Body)
  if err != nil {
    panic(err)
  }

  var u models.EditQuestion

  err = json.Unmarshal(body, &u)
  if err != nil {
    panic(err)
  }

  db, err := sql.Open("postgres", "password=password host=localhost dbname=online_test_dev sslmode=disable")
  if err != nil {
    panic(err)
  }
  questions, err:= db.Exec("CREATE TABLE IF NOT EXISTS questions(id int, title text, option_1 varchar(100), option_2 varchar(100), option_3 varchar(100), option_4 varchar(100), answer varchar(100), section_id int, CONSTRAINT section_id_key FOREIGN KEY(section_id) REFERENCES sections (id), PRIMARY KEY(id))")
  if err != nil || questions == nil {
    panic(err)
  }

  update_question, err := db.Query("UPDATE questions SET title = $1, option_1 = $2, option_2 = $3, option_3 = $4, option_4 = $5, answer = $6 where id = $7", u.Question, u.OptionA, u.OptionB, u.OptionC, u.OptionD, u.Answer, u.Id)
  if err != nil || update_question == nil {
    panic(err)
  }
  defer update_question.Close()

  b, err := json.Marshal(models.UpdateQuestionMessage{
    Success: "true",
    Message: "Question updated Successfully!",
    })
  if err != nil {
    panic(err)
  }
  rw.Header().Set("Content-Type", "application/json")
  rw.Write(b)
}


func (q questionController) DeleteQuestions(rw http.ResponseWriter, req *http.Request) {

  vars := mux.Vars(req)
  question_id := vars["id"]

  db, err := sql.Open("postgres", "password=password host=localhost dbname=online_test_dev sslmode=disable")
  if err != nil {
    panic(err)
  }
  questions, err:= db.Exec("CREATE TABLE IF NOT EXISTS questions(id int, title text, option_1 varchar(100), option_2 varchar(100), option_3 varchar(100), option_4 varchar(100), answer varchar(100), section_id int, CONSTRAINT section_id_key FOREIGN KEY(section_id) REFERENCES sections (id), PRIMARY KEY(id))")
  if err != nil || questions == nil {
    panic(err)
  }

  delete_question, err := db.Query("DELETE from questions where id = $1",question_id)
  if err !=nil {
    panic(err)
  }
  defer delete_question.Close()

  b, err := json.Marshal(models.UpdateQuestionMessage{
    Success: "true",
    Message: "Question deleted Successfully!",
    })
  if err != nil {
    panic(err)
  }
  rw.Header().Set("Content-Type", "application/json")
  rw.Write(b)

}