package account

import (
"database/sql"
"encoding/json"
"fmt"
"strconv"
"time"
"crypto/md5"
"io"
"github.com/Kedarnag13/Online_test/api/v1/models"
"github.com/asaskevich/govalidator"
"github.com/Kedarnag13/Online_test/api/v1/controllers"
_ "github.com/lib/pq"
"io/ioutil"
"log"
"net/http"
"regexp"
"encoding/hex"
)

type registrationController struct{}

var Registration registrationController

func (r registrationController) Create(rw http.ResponseWriter, req *http.Request) {
	body, err := ioutil.ReadAll(req.Body)
	flag := 1
	var u models.Register

	if err != nil {
		panic(err)
	}
	err = json.Unmarshal(body, &u)
	if err != nil {
		panic(err)
	}

	db, err := sql.Open("postgres", "password=password host=localhost dbname=online_test_dev sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}

	res, err := db.Query("SELECT email FROM users ")
	if err != nil {
		log.Fatal(err)
	}
	defer res.Close()

	fetch_id, err := db.Query("SELECT coalesce(max(id), 0) FROM users")
	if err != nil {
		log.Fatal(err)
	}
	defer fetch_id.Close()

	if flag == 1 {
		email := `^[a-z0-9._%+\-]+@[a-z0-9.\-]+\.[a-z]{2,4}$`
		exp, err := regexp.Compile(email)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println("phone_number:",u.Phone_number)
		if u.First_name == "" || u.Last_name == "" || u.Email == "" || !exp.MatchString(u.Email) || u.Password == "" || u.Password_confirmation == "" || u.College == "" || u.Year_of_passing == "" {

			_, err := govalidator.ValidateStruct(u)
			if err != nil {
				println("error: " + err.Error())
			}

			flag = 0
			b, err := json.Marshal(models.ErrorMessage{
				Success: "false",
				Error:   err.Error(),
				})
			if err != nil {
				log.Fatal(err)
			}
			rw.Header().Set("Content-Type", "application/json")
			rw.Write(b)
			goto create_user_end
		} else if u.Password != u.Password_confirmation {
			flag = 0
			b, err := json.Marshal(models.EmailErrorMessage{
				Success: "false",
				Error:   "Password and confirm password do not match!",
				})
			if err != nil {
				log.Fatal(err)
			}
			rw.Header().Set("Content-Type", "application/json")
			rw.Write(b)
			goto create_user_end
		}
	}

	if flag == 1 {
		for res.Next() { // email already exist condition
			var email string
			err = res.Scan(&email)
			if err != nil {
				log.Fatal(err)
			}

			if email == u.Email {
				b, err := json.Marshal(models.ErrorMessage{
					Success: "false",
					Error:   "Email id already exist",
					})
				if err != nil {
					log.Fatal(err)
				}
				rw.Header().Set("Content-Type", "application/json")
				rw.Write(b)
				fmt.Println("Email id already exist")
				flag = 0
				goto create_user_end
			}
		}

		// password and confirm password does not match =====================

		if u.Password != u.Password_confirmation {
			b, err := json.Marshal(models.ErrorMessage{
				Success: "false",
				Error:   "Password and Password_confirmation do not match",
				})
			if err != nil {
				log.Fatal(err)
			}
			rw.Header().Set("Content-Type", "application/json")
			rw.Write(b)
			goto create_user_end
		}

		// Insert into users table ======================================

		for fetch_id.Next() {
			var id int
			err = fetch_id.Scan(&id)

			if err != nil {
				log.Fatal(err)
			}
			id = id + 1

			var sStmt string = "insert into users (id, first_name, last_name, email, college, branch, phone_number, year_of_passing, password, password_confirmation) values ($1,$2,$3,$4,$5,$6,$7,$8,$9,$10)"
			db, err := sql.Open("postgres", "password=password host=localhost dbname=online_test_dev sslmode=disable")
			if err != nil {
				log.Fatal(err)
			}
			stmt, err := db.Prepare(sStmt)
			if err != nil {
				log.Fatal(err)
			}
			defer stmt.Close()


			key := []byte("traveling is fun")
			password := []byte(u.Password)
			confirm_password := []byte(u.Password_confirmation)
			encrypt_password := controllers.Encrypt(key, password)
			encrypt_password_confirmation := controllers.Encrypt(key, confirm_password)

			user_res, err := stmt.Exec(id, u.First_name, u.Last_name, u.Email, u.College, u.Branch, u.Phone_number, u.Year_of_passing, encrypt_password, encrypt_password_confirmation)
			if err != nil || user_res == nil {
				log.Fatal(err)
			}
			defer stmt.Close()

			// Create Session for the User =========================================

			auth_string := strconv.FormatInt(time.Now().Unix(), 10)
			h := md5.New()
			io.WriteString(h, auth_string)
			auth_token := hex.EncodeToString(h.Sum(nil))
			var session string = "insert into sessions (start_time, user_id, auth_token) values ($1,$2,$3)"
			ses, err := db.Prepare(session)
			if err != nil {
				log.Fatal(err)
			}
			start_time := time.Now()
			session_res, err := ses.Exec(start_time, id, string(auth_token))
			if err != nil || session_res == nil {
				log.Fatal(err)
			}

			fmt.Printf("StartTime: %v\n", time.Now())
			fmt.Println("User created Successfully!")

			user := models.Register{id, u.First_name, u.Last_name, u.Email, u.Password, u.Password_confirmation, u.College, u.Branch, u.Year_of_passing, u.Phone_number}

			b, err := json.Marshal(models.SignUp{
				Success: "true",
				Message: "User created Successfully!",
				User:    user,
				Session: models.Session{id, start_time, string(auth_token)},
				})

			if err != nil || res == nil {
				log.Fatal(err)
			}

			rw.Header().Set("Content-Type", "application/json")
			rw.Write(b)
		}
		// defer fetch_id.Close()
	}
	create_user_end:
	db.Close()
}
