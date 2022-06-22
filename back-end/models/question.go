package models

type Choice struct {
	ChoiceId string `json:"choiceId" bson:"choiceId"`
	Content  string `json:"content" bson:"content"`
}

type Question struct {
	Type        string   `json:"questionType" bson:"questionType"`
	QuestionId  int      `json:"questionId" bson:"questionId"`
	Description string   `json:"description" bson:"description"`
	Choices     []Choice `json:"choices" bson:"choices"`
	Answer      string   `json:"answer" bson:"answer"`
}

type Header struct {
	HeaderId    int        `json:"headerId" bson:"headerId"`
	Description string     `json:"description" bson:"description"`
	Image       string     `json:"image" bson:"image"`
	Questions   []Question `json:"questions" bson:"questions"`
}

type Container struct {
	Headers []Header `json:"headers"`
}

type Answer struct {
	HeaderId  int      `json:"headerId"`
	Solutions []string `json:"solutions"`
	Point     []int    `json:"point"`
}

type AnsContainer struct {
	Answers []Answer `json:"answers"`
}
