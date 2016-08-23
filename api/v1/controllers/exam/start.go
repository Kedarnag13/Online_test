package exam

import (
"database/sql"
"encoding/json"
"github.com/Kedarnag13/Online_test/api/v1/models"
_ "github.com/lib/pq"
"net/http"
"github.com/gorilla/mux"
"log"
)

type examController struct{}

var Exam examController

func (e examController) Questions(rw http.ResponseWriter, req *http.Request) {

  vars := mux.Vars(req)
  section_name := vars["section_name"]

  log.Printf("section_name: %v",section_name)
  db, err := sql.Open("postgres", "password=password host=localhost dbname=online_test_dev sslmode=disable")
  if err != nil {
    panic(err)
  }
  questions, err:= db.Exec("CREATE TABLE IF NOT EXISTS questions(id int, title text, option_1 varchar(100), option_2 varchar(100), option_3 varchar(100), option_4 varchar(100), answer varchar(100), section_id int, CONSTRAINT section_id_key FOREIGN KEY(section_id) REFERENCES sections (id), PRIMARY KEY(id))")
  if err != nil || questions == nil {
    panic(err)
  }
  get_section_id, err := db.Query("SELECT id from sections where name = $1", section_name)
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

  get_questions, err := db.Query("SELECT id, title, option_1, option_2, option_3, option_4, image, image_url FROM questions WHERE section_id=$1 GROUP BY id ORDER BY random() LIMIT 20", section_id)
  if err != nil || get_questions == nil {
    panic(err)
  }
  defer get_questions.Close()
  db.Close()

  questions_section := []models.Question{}

  for get_questions.Next() {
    var id int
    var title string
    var option_1 string
    var option_2 string
    var option_3 string
    var option_4 string
    var options []string
    var image string
    var image_url string

    var question_details models.Question

    err := get_questions.Scan(&id, &title, &option_1, &option_2, &option_3, &option_4, &image, &image_url)
    if err != nil {
      panic(err)
    }
    log.Printf("fetch Question_id:%v",id)
    options = append(options, option_1)
    options = append(options, option_2)
    options = append(options, option_3)
    options = append(options, option_4)

    if image != "nil" && image_url == "nil"{
      image_url = Fetch_image_url(image)
    }
    question_details = models.Question{id, title, image_url, options}
    questions_section = append(questions_section, question_details)
  }
  b, err := json.Marshal(models.QuestionResponseMessage{
    Success:     "true",
    Message:     "Questions per section",
    SectionName:  section_name,
    QuestionList: questions_section,
    })
  if err != nil {
    panic(err)
  }
  rw.Header().Set("Content-Type", "application/json")
  rw.Write(b)
  db.Close()
}
