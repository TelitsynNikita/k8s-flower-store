package models

type LogoPass struct {
	Login    string `json:"login" validate:"required,gt=4,lte=40"`
	Password string `json:"password" validate:"required,gt=8,lte=40"`
}
