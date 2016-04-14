package models

import (
	"time"
)

// Registration struct [account/sign_up]
type Register struct {
	Id                    int    `valid:"numeric"`
	First_name            string `valid:"alphanum,required"`
	Last_name             string `valid:"alphanum,required"`
	Email                 string `valid:"email,required"`
	Password              string `valid:"alphanum,required"`
	Password_confirmation string `valid:"alphanum,required"`
	College								string `valid:"alphanum,required"`
	Branch                string `valid:"alphanum",required`
	Year_of_passing				string `valid:"alphanum",required`
	City									string `valid:"alphanum",required`
	Phone_number					string `required`
	Batch									string `valid:"alphanum",required`
}


type ErrorMessage struct {
	Success string
	Error   string
}

type EmailPasswordErrorMessage struct {
	Success string
	Email_error string
	Phone_number_error string
}

type FieldErrorMessage struct {
	Success string
	Error   []string
}

type SignUp struct {
	Success string
	Message string
	User    Register
	Session Session
}


// Session struct [account/session]
type Session struct {
	SessionId int
	StartTime time.Time
	Auth_token string
}

// Sign_up struct end

type UserDetails struct {
	Id                 int
	Firstname          string
	Lastname           string
	Email              string
	User_thumbnail     string
	User_thumbnail_web string
}

type InviteEmail struct {
	SenderId      int    `valid:"numeric,required"`
	RecieverEmail string `valid:"email,required"`
}

// Log_in struct

type LogIn struct {
	Phone_number string `valid:"required"`
	Password string `valid:"alphanum,required"`
}

type LogOut struct {
	Success string
	Message string
}

type SuccessfulLogIn struct {
	Success string
	Message string
	User_id int
	User_role string
	Session Session
}


// Message struct [controllers/account]
// Common for sign_up, session and password
type Message struct {
	Success string
	Message string
	User    Register
}


type EmailMessage struct {
	Success string
	Message string
	User    InviteEmail
}

type PasswordErrorMessage struct {
	Success string
	Password_error  string
}

type EmailErrorMessage struct {
	Success string
	Email_error   string
}

type PhoneNumberErrorMessage struct {
	Success string
	Phone_number_error   string
}

type AdminSuccessMessage struct {
	Success string
	Message string
}
