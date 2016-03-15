package models


// type Logical struct {
// 	Data string
// }
//
// type Aptitude struct {
// 	Data string
// }

type Question struct {
	Id int
	Title string
	Option_1 string
	Option_2 string
	Option_3 string
	Option_4 string
	Answer string
}

type QuestionResponseMessage struct {
	Success string
	Message string
	QuestionList []Question
}
type QuestionResponse struct {
	SectionId int
	Questions []Answer
}

type Answer struct {
	QuestionId int
	AnswerOption string
}


type Result struct {
	Section int
	TotalQuestions int
	Score int
}
