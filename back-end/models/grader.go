package models

type Grader_Sub_Questions map[string][]string

type GraderQuestions struct {
	TestAutograder Grader_Sub_Questions `json:"test_autograder"`
}

type GraderTest struct {
	Answers   GraderQuestions `json:"answers"`
	Solutions GraderQuestions `json:"solutions"`
}

type GraderTestError struct {
	Type    string `json:"type" example:"Grader"`
	Message string `json:"message" example:"<---Running--->\n<---Failure--->\nOutput from grader..."`
}

type GraderTestSuccess struct {
	Status int    `json:"status" example:"201"`
	Type   int    `json:"type" example:"0"`
	Error  any    `json:"error"`
	Data   string `json:"data" example:"<---Running--->\n<---Success--->\nOutput from grader..."`
}
