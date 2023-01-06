package controllers

import (
	"github.com/go_crud_api/interfaces"
)

type Resources struct {
	db interfaces.DB
}

// NewConfig is a constructor for the Resources
func NewConfig(db interfaces.DB) *Resources {
	return &Resources{db}
}
