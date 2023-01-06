package models

import (
	"gorm.io/gorm"

	"github.com/go_crud_api/interfaces"
)

// User es el modelo para la entidad User en la base de datos
type User struct {
	gorm.Model
	Name  string `json:"name"`
	Email string `json:"email"`
	Phone string `json:"phone"`
}

// create user in the database
func (user *User) Create(db interfaces.DB) (*User, error) {
	if err := db.Create(&user).Error; err != nil {
		return nil, err
	}
	return user, nil
}

// edit user in the database
func (user *User) Edit(db interfaces.DB, id interface{}) (*User, error) {

	db.First(&user, user.ID)
	if err := db.Save(&user).Error; err != nil {
		return nil, err
	}
	return user, nil
}

// delete user in the database
func (user *User) Delete(db interfaces.DB, id interface{}) (*User, error) {
	if err := db.Delete(&user, id).Error; err != nil {
		return nil, err
	}
	return user, nil
}

// get one user from the database
func (user *User) Get(db interfaces.DB, id interface{}) (*User, error) {
	if err := db.First(&user, id).Error; err != nil {
		return nil, err
	}
	return user, nil
}
