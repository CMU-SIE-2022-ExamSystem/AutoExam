package models

type User_Info struct {
	Email      string `json:"email"`
	First_name string `json:"first_name"`
	Last_name  string `json:"last_name"`
	School     string `json:"school"`
	Major      string `json:"major"`
	Year       string `json:"year"`
}

type User_Courses struct {
	Name         string `json:"name"`
	Semester     string `json:"semester"`
	Late_slack   int64  `json:"late_slack"`
	Grace_days   int64  `json:"grace_days"`
	Display_name string `json:"display_name"`
	Auth_level   string `json:"auth_level"`
}

type Autolab_Info_Front struct {
	Scope     string `json:"scope"`
	Client_id string `json:"clientId"`
}

type Assessments struct {
	Name             string `json:"name"`
	Display_name     string `json:"display_name"`
	Start_at         string `json:"start_at"`
	Due_at           string `json:"due_at"`
	End_at           string `json:"end_at"`
	Category_name    string `json:"category_name"`
	Grading_deadline string `json:"grading_deadline"`
}

type Submissions struct {
	Version    int         `json:"version"`
	Filename   string      `json:"filename"`
	Created_at string      `json:"created_at"`
	Scores     interface{} `json:"scores"`
}

type Submit struct {
	Version  int    `json:"version"`
	Filename string `json:"filename"`
}

type Course_User_Data struct {
	First_name   string `json:"first_name"`
	Last_name    string `json:"last_name"`
	Email        string `json:"email"`
	School       string `json:"school"`
	Major        string `json:"major"`
	Year         string `json:"year"`
	Lecture      string `json:"lecture"`
	Section      string `json:"section"`
	Grade_policy string `json:"grade_policy"`
	Nickname     string `json:"nickname"`
	Dropped      bool   `json:"dropped"`
	Auth_level   string `json:"auth_level"`
}

type Course_User_err struct {
	Error string `json:"error"`
}
