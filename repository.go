package taskService

import "gorm.io/gorm"

// Основные методы CRUD - Create, Read, Update, Delete

type TaskRepository interface {
	Create(task Task) (Task, error)
	GetTaskByID(id string) (Task, error)
	Update(id string, task Task) (Task, error)
	DeleteTask(id string) error
}

type tskRepository struct {
	db *gorm.DB
}

func NewTaskRepository(db *gorm.DB) TaskRepository {
	return &tskRepository{db: db}
}

func (r *tskRepository) Create(task Task) (Task, error) {
	if err := r.db.Create(&task).Error; err != nil {
		return Task{}, err
	}
	return task, nil
}

func (r *tskRepository) GetTaskByID(id string) (Task, error) {
	var tsk Task
	err := r.db.First(&tsk, "id = ?", id).Error
	return tsk, err
}

func (r *tskRepository) Update(id string, task Task) (Task, error) {
	result := r.db.Model(&Task{}).Where("id = ?", id).Updates(task)
	if result.Error != nil {
		return Task{}, result.Error
	}
	// Получите обновленную задачу
	var updatedTask Task
	if err := r.db.First(&updatedTask, "id = ?", id).Error; err != nil {
		return Task{}, err
	}
	return updatedTask, nil
}

func (r *tskRepository) DeleteTask(id string) error {
	return r.db.Delete(&Task{}, id).Error
}
