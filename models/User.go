package models

import "github.com/jinzhu/gorm"

// User es el modelo para la entidad User en la base de datos
type User struct {
	gorm.Model
	Name  string `json:"name"`
	Email string `json:"email"`
	Phone string `json:"phone"`
}
