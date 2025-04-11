package handlers

import (
	"RestApi/internal/taskService"
	"net/http"

	"github.com/labstack/echo/v4"
)

type TaskHandler struct {
	service taskService.TaskService
}

func NewTaskHAndler(s taskService.TaskService) *TaskHandler {
	return &TaskHandler{service: s}
}

func (h *TaskHandler) GetTask(c echo.Context) error {
	id := c.Param("id")
	tasks, err := h.service.GetTaskByID(id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Couldn't get the tasks"})
	}
	return c.JSON(http.StatusOK, tasks)
}

func (h *TaskHandler) PostTask(c echo.Context) error {
	var task taskService.Task
	if err := c.Bind(&task); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Could not add task"})
	}

	createdTask, err := h.service.CreateTask(task)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Could not create task"})
	}
	return c.JSON(http.StatusCreated, createdTask)
}

func (h *TaskHandler) PatchTask(c echo.Context) error {
	id := c.Param("id")

	var updatedTask taskService.Task
	if err := c.Bind(&updatedTask); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid request data"})
	}

	// 2. Вызовите сервис с двумя аргументами
	result, err := h.service.UpdateTask(id, updatedTask)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, result)
}

func (h *TaskHandler) DeleteTask(c echo.Context) error {
	id := c.Param("id")

	if err := h.service.DeleteTask(id); err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "The task was not deleted"})
	}
	return c.NoContent(http.StatusNoContent)
}
