package usecase

import "todo-api/repository"

// Интерфейс TaskUsecase — то, что увидит handler.
// В нём 3 метода — по одному на каждую операцию, которую должен уметь usecase-слой
type TaskUsecase interface {
	CreateTask(title, description string) error// создание задачи
	GetTask(id int) (repository.Task, error)// получение задачи по id
	ListTask() ([]repository.Task, error)// получение всех задач
}
// taskUsecase — конкретная реализация интерфейса TaskUsecase
type taskUsecase struct {
	repo repository.TaskRepository
}

func NewTaskUsecase(repo repository.TaskRepository) *taskUsecase {
	return &taskUsecase{repo: repo}
}
func (u *taskUsecase) CreateTask(title, description string) error {
