package models

type CreateQuestion struct {
  Section string `valid:"alphanum,required"`
  Question string `valid:"duck,required"`
  OptionA string `valid:"duck,required"`
  OptionB string `valid:"duck,required"`
  OptionC string `valid:"duck,required"`
  OptionD string `valid:"duck,required"`
  Answer string `valid:"duck",required`
}


type EditQuestion struct {
  Section int 
  Question string `valid:"duck,required"`
  OptionA string `valid:"duck,required"`
  OptionB string `valid:"duck,required"`
  OptionC string `valid:"duck,required"`
  OptionD string `valid:"duck,required"`
  Answer string `valid:"duck",required`
  Id int
}

type CreateQuestionStatusMessage struct {
  Success string
  Message string
}

type UpdateQuestionMessage struct {
  Success string
  Message string
}

type FetchQuestion struct {
  Id int
  Title string
  Option_a string
  Option_b string
  Option_c string
  Option_d string
  Answer string
  Section_id int
}

type FetchQuestionResponseMessage struct {
  Success string
  Message string
  QuestionList []FetchQuestion
}