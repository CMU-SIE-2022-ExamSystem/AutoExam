package models

type User struct {
	ID            uint   `json:"id" gorm:"primaryKey"`
	Email         string `json:"email" gorm:"unique"`
	Access_token  string `json:"access_token"`
	Refresh_token string `json:"refresh_token"`
	First_name    string `json:"first_name"`
	Last_name     string `json:"last_name"`
	Create_at     int64  `json:"create_at"`
	Expires_in    int64  `json:"expires_in"`
}

type UserToken struct {
	ID    uint   `json:"id" gorm:"primaryKey"`
	Email string `json:"email"`
}

func (User) TableName() string {
	return "user"
}
