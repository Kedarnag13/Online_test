package feedback

import (
"log"
"net/http"
"io/ioutil"
"github.com/Kedarnag13/Online_test/api/v1/models"
"encoding/json"
"github.com/asaskevich/govalidator"
"database/sql"
) 

type feedbackController struct{}

var Feedback feedbackController

func (e feedbackController) Create(rw http.ResponseWriter, req *http.Request) { 
	body, err := ioutil.ReadAll(req.Body)

	var f models.Feedback

	if err != nil {
		panic(err)
	}
	err = json.Unmarshal(body, &f)
	if err != nil {
		panic(err)
	}

	db, err := sql.Open("postgres", "password=password host=localhost dbname=online_test_dev sslmode=disable")
	if err != nil {
		panic(err)
	}
	if f.Verbal_section == "" || f.Logical_section == "" || f.Aptitude_section == "" {
		_, err := govalidator.ValidateStruct(f)
		if err != nil {
			println("error: " + err.Error())
		}
		b, err := json.Marshal(models.ErrorMessage{
			Success: "false",
			Error:   err.Error(),
			})
		if err != nil {
			panic(err)
		}
		rw.Header().Set("Content-Type", "application/json")
		rw.Write(b)
		goto feedback
	} else {
		var feedback_sStmt string = "insert into feedbacks (verbal, logical, aptitude, description) values ($1, $2, $3, $4)"
		feedback_prepare_stmt, err := db.Prepare(feedback_sStmt)
		if err != nil || feedback_prepare_stmt == nil {
			panic(err)
		}
		defer feedback_prepare_stmt.Close()
		feedback_exec, err := 	feedback_prepare_stmt.Exec(f.Verbal_section, f.Logical_section, f.Aptitude_section, f.Description)
		if err != nil || feedback_exec == nil {
			panic(err)
		}
		b, err := json.Marshal(models.FeedbackResponse{
			Success: "true",
			Message: "Feedback recorded Successfully!",
			})

		if err != nil || feedback_exec == nil {
			log.Fatal(err)
		}

		rw.Header().Set("Content-Type", "application/json")
		rw.Write(b)
	}
	feedback:
	db.Close()
}