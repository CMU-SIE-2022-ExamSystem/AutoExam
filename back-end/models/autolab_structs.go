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
