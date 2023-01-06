package models_test

import (
	"testing"

	"github.com/go_crud_api/mocks"
	"github.com/go_crud_api/models"
	"github.com/stretchr/testify/assert"
)

var (
	userFull = models.User{
		Name:  "Name",
		Email: "email@domain.com",
		Phone: "1234567890",
	}
)

// test the creation of a user
func TestUserCreation(t *testing.T) {

	// create mock db
	dbMock := &mocks.MockDB{}

	_, err := userFull.Create(dbMock)

	if err != nil {
		assert.Fail(t, "Error creating user")
	}

	assert.Equal(t, nil, err)

}

// test the edition of a user
func TestUserEdition(t *testing.T) {

	// create mock db
	dbMock := &mocks.MockDB{}
	id := "1"

	_, err := userFull.Edit(dbMock, id)

	if err != nil {
		assert.Fail(t, "Error editing user")
	}

	assert.Equal(t, nil, err)

}

// test the deletion of a user
func TestUserDeletion(t *testing.T) {

	// create mock db
	dbMock := &mocks.MockDB{}
	id := "1"

	_, err := userFull.Delete(dbMock, id)

	if err != nil {
		assert.Fail(t, "Error deleting user")
	}

	assert.Equal(t, nil, err)

}

// test the get of a user
func TestUserGet(t *testing.T) {

	// create mock db
	dbMock := &mocks.MockDB{}
	userID := "1"

	_, err := userFull.Get(dbMock, userID)

	if err != nil {
		assert.Fail(t, "Error getting user")
	}

	assert.Equal(t, nil, err)

}
