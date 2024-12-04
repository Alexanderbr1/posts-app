package domain

type User struct {
	Id       int    `json:"-" db:"id"`
	Name     string `json:"name" binding:"required,min=2" example:"Dmitry"`
	Username string `json:"username" binding:"required,min=2" example:"mdmitry"`
	Password string `json:"password" binding:"required,min=8" example:"12345678"`
}

type SignInInput struct {
	Username string `json:"username" binding:"required,min=2" example:"mdmitry"`
	Password string `json:"password" binding:"required,min=8" example:"12345678"`
}
