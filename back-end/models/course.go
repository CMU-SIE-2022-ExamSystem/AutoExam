package models

type Assessments struct {
	Name             string `json:"name"`
	Display_name     string `json:"display_name"`
	Start_at         string `json:"start_at"`
	Due_at           string `json:"due_at"`
	End_at           string `json:"end_at"`
	Category_name    string `json:"category_name"`
	Grading_deadline string `json:"grading_deadline"`
}
