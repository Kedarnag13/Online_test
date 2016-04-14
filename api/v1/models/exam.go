package models

import "time"

type Question struct {
	Id int
	Title string
	Options []string
}

type QuestionResponseMessage struct {
	Success string
	Message string
	SectionId int
	QuestionList []Question
}
type QuestionResponse struct {
	SectionId int
	UserId int
	Questions []Answer
}

type Answer struct {
	QuestionId int
	Answer string
}


type Result struct {
	Section int
	TotalQuestions int
	Score int
}

type UserResult struct {
	First_name string
	Last_name string
	Email string
	Phone_number string
	City string
	Batch string
	Section_1_score int
	Section_2_score int
	Section_3_score int
	Total_score int
	StartTime time.Time
	EndTime time.Time
}

type Report struct {
	Total_users int
	Report_list []UserResult
}
