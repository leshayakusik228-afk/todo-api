package usecase

import "todo-api/repository"

// Интерфейс TaskUsecase — то, что увидит handler.
// В нём 3 метода — по одному на каждую операцию, которую должен уметь usecase-слой
type TaskUsecase interface {
	CreateTask(title, description string) error // создание задачи
	GetTask(id int) (repository.Task, error)    // получение задачи по id
	ListTask() ([]repository.Task, error)       // получение всех задач
}

// taskUsecase — конкретная реализация интерфейса TaskUsecase
type taskUsecase struct {
	repo repository.TaskRepository
}

func NewTaskUsecase(repo1 repository.TaskRepository) *taskUsecase {
	return &taskUsecase{repo: repo1}

}

func (u *taskUsecase) ListTask() ([]repository.Task, error) {
	return u.repo.List() //вызываем метод List() у репозитория, который вернёт слайс задач и ошибку
}

func (u *taskUsecase) GetTask(id int) (repository.Task, error) {
	return u.repo.GetByID(id) //вызываем метод GetByID() у репозитория, который вернёт задачу и ошибку
}
func (u *taskUsecase) CreateTask(title, description string) error {
	//1. получить все задачичтобы понят сколько их
	tasks, err := u.repo.List() //получаем список всех задач, чтобы определить новый ID
	if err != nil {
		return err
	}
	//новая задача
	task := repository.Task{
		ID:          len(tasks) + 1, //новый ID будет на 1 больше чем количество задач
		Title:       title,
		Description: description,
		Status:      "new", //по умолчанию статус новой задачи "new"
	}
	return u.repo.Create(task) //вызываем метод Create() у репозитория, который создаст задачу и вернёт ошибку если что-то не так
}
