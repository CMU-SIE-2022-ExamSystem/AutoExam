package dao

///the version for storing, with answer
type Choice struct {
	ChoiceId string `json:"choiceId" bson:"choiceId"`
	Content  string `json:"content" bson:"content"`
}

// sub question with answer
type Question struct {
	Type        string   `json:"question_type" bson:"questionType"`
	QuestionId  int      `json:"question_id" bson:"questionId"`
	Description string   `json:"description" bson:"description"`
	Choices     []Choice `json:"choices" bson:"choices"`
	Answer      []string `json:"answer" bson:"answer"` //answer is changed into an array of strings
}

// question header with answer
type Question_Header struct {
	HeaderId    int        `json:"id" bson:"_id"`
	Tag         string     `json:"question_tag" bson:"questionTag"`
	Title       string     `json:"title" bson:"title"`
	Description string     `json:"description" bson:"description"` //html string
	Questions   []Question `json:"questions" bson:"questions"`
}

// container with answer
type Container struct {
	Data []Question_Header `json:"data"`
}

//the version for reading, no answer
type Choice1 struct {
	ChoiceId string `json:"choiceId" bson:"choiceId"`
	Content  string `json:"content" bson:"content"`
}

// sub question without answer
type Question_Without_Answer struct {
	Type        string    `json:"questionType" bson:"questionType"`
	QuestionId  int       `json:"questionId" bson:"questionId"`
	Description string    `json:"description" bson:"description"`
	Choices     []Choice1 `json:"choices" bson:"choices"`
}

// question header without answer
type Question_Header_Without_Answer struct {
	HeaderId    int                       `json:"id" bson:"_id"`
	Tag         string                    `json:"question_tag" bson:"questionTag"`
	Title       string                    `json:"title" bson:"title"`
	Description string                    `json:"description" bson:"description"` //html string
	Questions   []Question_Without_Answer `json:"questions" bson:"questions"`
}

type Container_Without_Answer struct {
	Data []Question_Header_Without_Answer `json:"data"`
}

type Tags_Return struct {
	Tags []string `yaml:"tags" json:"tags"`
}
