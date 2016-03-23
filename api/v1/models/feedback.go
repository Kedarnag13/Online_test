package models

type Feedback struct {
	Verbal_section string `valid:"alphanum,required"`
	Logical_section string `valid:"alphanum,required"`
	Aptitude_section string `valid:"alphanum,required"`
	Description string `valid:"duck"`
}

type FeedbackResponse struct {
	Success string
	Message string 
}