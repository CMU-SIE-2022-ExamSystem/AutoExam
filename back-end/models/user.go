package models

type User struct {
	ID            uint   `json:"id" gorm:"primaryKey"`
	Email         string `json:"email"`
	Access_token  string `json:"access_token"`
	Refresh_token string `json:"refresh_token"`
	First_name    string `json:"first_name"`
	Last_name     string `json:"last_name"`
}

type UserToken struct {
	ID    uint   `json:"id" gorm:"primaryKey"`
	Token string `json:"token"`
}

func (User) TableName() string {
	return "user"
}
