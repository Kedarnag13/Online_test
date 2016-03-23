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
	Options []string
}

type QuestionResponseMessage struct {
	Success string
	Message string
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
