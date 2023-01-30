package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
	"go-fiber-gorm/internal/database"
	"go-fiber-gorm/router"
)

func main() {

	// INITIAL DATABASE
	db, err := database.ConnectDB()
	if err != nil {
		logrus.Fatal(err.Error())
	}

	defer db.Clauses()

	app := fiber.New()

	// INITIAL ROUTER
	router.RouterInit(app)

	app.Listen(":8080")
}
