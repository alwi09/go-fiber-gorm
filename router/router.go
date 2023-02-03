package router

import (
	"github.com/gofiber/fiber/v2"
	"go-fiber-gorm/internal/handler"
	"gorm.io/gorm"
)

func RouterInit(router *fiber.App, db *gorm.DB) error {

	handler := handler.NewTodolistHandlerImpl(db)

	router.Get("/api/GetAllTodo", handler.GetAllTodo)

	return nil
}
