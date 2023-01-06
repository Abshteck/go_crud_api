package mocks

import (
	"gorm.io/gorm"
)

// mockDB is a mock of the DB interface
type MockDB struct{}

// Create is a mock of the Create method
func (m *MockDB) Create(value interface{}) *gorm.DB {
	return &gorm.DB{}
}

// Edit is a mock of the Edit method
func (m *MockDB) Save(value interface{}) *gorm.DB {
	return &gorm.DB{}
}

// Delete is a mock of the Delete method
func (m *MockDB) Delete(value interface{}, conds ...interface{}) *gorm.DB {
	return &gorm.DB{}
}

// Get is a mock of the Get method
func (m *MockDB) First(out interface{}, where ...interface{}) *gorm.DB {
	return &gorm.DB{}
}
