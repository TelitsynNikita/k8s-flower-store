package models

type Flower struct {
	FlowerID    uint64 `json:"flower_id" validate:"omitempty,gt=0,lte=18446744073709551615"`
	FlowerName  string `json:"flower_name" validate:"required,gt=0,lte=20"`
	FlowerPrice uint64 `json:"flower_price" validate:"required,gt=0,lte=100000"`
}
