package interfaces

import (
	"gorm.io/gorm"
)

// DB is the interface for the database
type DB interface {
	Create(value interface{}) *gorm.DB
	Save(value interface{}) *gorm.DB
	Delete(value interface{}, conds ...interface{}) *gorm.DB
	First(out interface{}, where ...interface{}) *gorm.DB
}
