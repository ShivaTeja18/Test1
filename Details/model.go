package Details

import "gorm.io/gorm"

type EMP struct {
	gorm.Model
	Name     string `json:"name"`
	Password string `json:"password"`
	City     string `json:"city"`
}

type Handler struct {
	DB *gorm.DB
}
