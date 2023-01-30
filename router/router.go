package router

import (
	"github.com/gofiber/fiber/v2"
	"go-fiber-gorm/internal/handler"
)

func RouterInit(router *fiber.App) {

	router.Get("/GetAllTodo", handler.GetAllTodo)

}
