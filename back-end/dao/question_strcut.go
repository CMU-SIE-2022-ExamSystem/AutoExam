package dao
///the version for storing, with answer
type Choice struct {
	ChoiceId string `json:"choiceId" bson:"choiceId"`
	Content string `json:"content" bson:"content"`
}

type Question struct {
	Type string `json:"questionType" bson:"questionType"`
	QuestionId int `json:"questionId" bson:"questionId"`
	Description string `json:"description" bson:"description"`
	Choices []Choice `json:"choices" bson:"choices"`
	Answer []string `json:"answer" bson:"answer"` //answer is changed into an array of strings
}

type Header struct {
	HeaderId int `json:"headerId" bson:"headerId"`
	Tag string `json:"questionTag" bson:"questionTag"`
	Description string `json:"description" bson:"description"` //html string
	Questions []Question `json:"questions" bson:"questions"`
}

type Container struct {
	Data []Header `json:"data"`
}



//the version for reading, no answer
type Choice1 struct {
	ChoiceId string `json:"choiceId" bson:"choiceId"`
	Content string `json:"content" bson:"content"`
}

type Question1 struct {
	Type string `json:"questionType" bson:"questionType"`
	QuestionId int `json:"questionId" bson:"questionId"`
	Description string `json:"description" bson:"description"`
	Choices []Choice1 `json:"choices" bson:"choices"`
}

type Header1 struct {
	HeaderId int `json:"headerId" bson:"headerId"`
	Tag string `json:"questionTag" bson:"questionTag"`
	Description string `json:"description" bson:"description"` //html string
	Questions []Question1 `json:"questions" bson:"questions"`
}

type Container1 struct {
	Data []Header1 `json:"data"`
}