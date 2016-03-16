package models

type Question struct {
	Id int
	Title string
	Option_1 string
	Option_2 string
	Option_3 string
	Option_4 string
	Answer string
}

type QuestionResponse struct {
	Success string
	Message string
	QuestionList []Question
}
