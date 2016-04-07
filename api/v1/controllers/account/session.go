package account

import (
  "database/sql"
  "encoding/json"
  "github.com/Kedarnag13/Online_test/api/v1/controllers"
  "github.com/Kedarnag13/Online_test/api/v1/models"
  "github.com/asaskevich/govalidator"
  _ "github.com/lib/pq"
  "io/ioutil"
  "io"
  "strconv"
  "time"
  "crypto/md5"
  "net/http"
  "encoding/hex"
  "fmt"
)

type sessionController struct{}

var Session sessionController

func (s sessionController) Create(rw http.ResponseWriter, req *http.Request) {

  body, err := ioutil.ReadAll(req.Body)
  var l models.LogIn

  if err != nil {
    panic(err)
  }

  flag := 0

  err = json.Unmarshal(body, &l)

  if l.Phone_number == "" || l.Password == ""  {

    _, err := govalidator.ValidateStruct(l)
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

    goto user_login_end
    }else {
      db, err := sql.Open("postgres", "password=password host=localhost dbname=online_test_dev sslmode=disable")
      if err != nil {
        panic(err)
      }
      get_user_id, err := db.Query("SELECT id, first_name, last_name, email, phone_number, password FROM users WHERE phone_number=$1", l.Phone_number)
      if err != nil {
        panic(err)
      }
      defer get_user_id.Close()


      for get_user_id.Next() {

        flag = 1
        var id int
        var first_name string
        var last_name string
        var email string
        var db_password string
        var phone_number string

        err := get_user_id.Scan(&id, &first_name, &last_name, &email, &phone_number, &db_password)
        if err != nil {
          panic(err)
        }

        check_session, err := db.Query("SELECT user_id from sessions where user_id = $1", id)
        if err !=nil {
          panic(err)
        }
        defer check_session.Close()

        for check_session.Next(){
          flag = 0
          var session_id int
          err := check_session.Scan(&session_id)
          if err !=nil {
            panic(err)
          }

          b, err := json.Marshal(models.ErrorMessage{
            Success: "false",
            Error:   "Session already exist",
          })

          if err != nil {
            panic(err)
          }
          rw.Header().Set("Content-Type", "application/json")
          rw.Write(b)

          goto user_login_end
        }


        key := []byte("traveling is fun")

        decrypt_password := controllers.Decrypt(key, db_password)

        if decrypt_password == l.Password {

          auth_string := strconv.FormatInt(time.Now().Unix(), 10)
          h := md5.New()
          io.WriteString(h, auth_string)
          auth_token := hex.EncodeToString(h.Sum(nil))
          var session string = "insert into sessions (start_time, user_id, auth_token) values ($1,$2,$3)"
          ses, err := db.Prepare(session)
          if err != nil {
            panic(err)
          }
          defer ses.Close()

          start_time := time.Now()
          session_res, err := ses.Exec(start_time, id, string(auth_token))
          if err != nil || session_res == nil {
            panic(err)
          }

          fmt.Printf("StartTime: %v\n", time.Now())
          fmt.Println("User Logged in Successfully!")

          b, err := json.Marshal(models.SuccessfulLogIn{
            Success: "true",
            Message: "User created Successfully!",
            User_id: id,
            Session: models.Session{id, start_time, string(auth_token)},
          })

          if err != nil {
            panic(err)
          }
          rw.Header().Set("Content-Type", "application/json")
          rw.Write(b)

          }else {
            b, err := json.Marshal(models.ErrorMessage{
              Success: "false",
              Error:   "Password does not match",
            })

            if err != nil {
              panic(err)
            }
            rw.Header().Set("Content-Type", "application/json")
            rw.Write(b)
          }

          goto user_login_end
        }

        if flag == 0 {
          b, err := json.Marshal(models.ErrorMessage{
            Success: "false",
            Error:   "Mobile number does not exist",
          })

          if err != nil {
            panic(err)
          }
          rw.Header().Set("Content-Type", "application/json")
          rw.Write(b)
        }
        db.Close()
      }

      user_login_end:
    }

func (s sessionController) Destroy(rw http.ResponseWriter, req *http.Request) {

  vars := mux.Vars(req)
  auth_token := vars["auth_token"]

  db, err := sql.Open("postgres", "password=password host=localhost dbname=online_test_dev sslmode=disable")
	if err != nil {
		panic(err)
	}

  delete_session, err := db.Query("DELETE FROM SESSIONS WHERE auth_token=$1", auth_token)
  if err != nil || delete_session == "" {
    panic(err)
  }
  b, err := json.Marshal(models.ErrorMessage{
    Success: "true",
    Error:   "Session destroyed successfully.",
  })

  if err != nil {
    panic(err)
  }
  rw.Header().Set("Content-Type", "application/json")
  rw.Write(b)
}
