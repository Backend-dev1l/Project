package main

import (
	"RestApi/internal/db"
	"RestApi/internal/handlers"
	"RestApi/internal/taskService"

	"log"

	"github.com/labstack/echo/v4"
)

// Основные методы ORM - Find, Create, Update, Delete

func main() {
	database, err := db.InitDB()
	if err != nil {
		log.Fatalf("Could not connect to DB: %v", err)
	}
	e := echo.New()

	tskRepo := taskService.NewTaskRepository(database)
	tskService := taskService.NewTaskService(tskRepo)
	tskHandlers := handlers.NewTaskHAndler(tskService)

	e.DELETE("/tasks/:id", tskHandlers.DeleteTask)
	e.GET("/tasks", tskHandlers.GetTask)
	e.POST("/tasks", tskHandlers.PostTask)
	e.PATCH("/tasks/:id", tskHandlers.PatchTask)
	e.Start(":8081")
}
