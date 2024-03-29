package models

type Course_Info_Front struct {
	Name         string `json:"name"`
	Display_name string `json:"display_name"`
	Auth_level   string `json:"auth_level"`
	Base_course  string `json:"base_course"`
}

type Course struct {
	ID   uint   `json:"id" gorm:"primaryKey"`
	Name string `json:"name" gorm:"unique"`
}

type Base_Course_Relationship struct {
	Course_name string `json:"course_name" gorm:"unique"`
	Base_course string `json:"base_course"`
}

func (Course) TableName() string {
	return "course"
}

type Basecourse struct {
	Name string `json:"name"`
}

type Download_Answer struct {
	Eamil string `json:"email"`
}
