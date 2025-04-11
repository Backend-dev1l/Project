package taskService

type TaskService interface {
	CreateTask(task Task) (Task, error)
	GetTaskByID(id string) (Task, error)
	UpdateTask(id string, updatedTask Task) (Task, error)
	DeleteTask(id string) error
}

type tskService struct {
	repo TaskRepository
}

func NewTaskService(r TaskRepository) TaskService {
	return &tskService{repo: r}
}

func (s *tskService) GetTask() (Task, error)

// CreateTask implements TaskService.
func (s *tskService) CreateTask(task Task) (Task, error) { // Добавьте параметр
	// Реализация метода
	createdTask, err := s.repo.Create(task)
	if err != nil {
		return Task{}, err
	}
	return createdTask, nil
}

// DeleteTask implements TaskService.
func (s *tskService) DeleteTask(id string) error {
	return s.repo.DeleteTask(id)
}

// GetTaskByID implements TaskService.
func (s *tskService) GetTaskByID(id string) (Task, error) {
	return s.repo.GetTaskByID(id)
}

// UpdateTask implements TaskService.
func (s *tskService) UpdateTask(id string, updatedTask Task) (Task, error) {
	// 1. Получите существующую задачу
	existingTask, err := s.repo.GetByID(id)
	if err != nil {
		return Task{}, err
	}

	// 3. Сохраните изменения
	return s.repo.Update(id, existingTask)

}
