package controllers_test

import (
	"bytes"
	"encoding/json"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/go_crud_api/controllers"
	"github.com/go_crud_api/mocks"
	"github.com/stretchr/testify/assert"
)

type userCreationCorrectRequest struct {
	Name  string `json:"name"`
	Email string `json:"email"`
	Phone string `json:"phone"`
}

// test CreateUser method
func TestUserCreation(t *testing.T) {
	// set gin to test mode and create router
	gin.SetMode(gin.TestMode)
	router := gin.New()

	// create mock db and resources
	dbMock := &mocks.MockDB{}
	res := controllers.NewConfig(dbMock)

	// create request body for user creation
	request := userCreationCorrectRequest{
		Name:  "Emilio",
		Email: "emilio.silva.dev@gmail.com",
		Phone: "1234567890",
	}

	var buf bytes.Buffer
	err := json.NewEncoder(&buf).Encode(request)
	if err != nil {
		t.Fatal(err)
	}

	// attach function to router
	router.POST("/user", res.CreateUser)

	// create gin context with request body
	c, _ := gin.CreateTestContext(httptest.NewRecorder())
	c.Request = httptest.NewRequest("POST", "/user", &buf)

	// call function
	res.CreateUser(c)

	// check response
	assert.Equal(t, 200, c.Writer.Status())
}
