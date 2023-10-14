package models

type User struct {
	Id       uint   `gorm:"primaryKey" json:"id" form:"id"`
	Name     string `gorm:"type:varchar(50);not null" json:"name"`
	Email    string `gorm:"type:varchar(50);not null;unique" json:"email"`
	Password string `gorm:"not null" json:"password"`
}

type RegisterResponse struct {
	Id    uint   `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

type LoginResponse struct {
	Id    uint   `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
	Token string `json:"token"`
}
