package models

type User struct {
	Name     string `json:"name" validate:"required,gt=1,lte=20"`
	LastName string `json:"last_name" validate:"required,gt=1,lte=20"`
	LogoPass
}
