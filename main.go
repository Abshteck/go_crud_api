package main

import (
	"github.com/go_crud_api/routes"
	"github.com/go_crud_api/utils"

	"gorm.io/gorm"
)

var (
	db *gorm.DB = utils.ConnectDB()
)

func main() {
	defer utils.DisconnectDB(db)

	//run all routes
	routes.Routes()
}
