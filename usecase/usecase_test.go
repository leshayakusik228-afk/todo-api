package usecase

import (
	"testing"
	"todo-api/repository"
)

func TestCreateTask(t *testing.T) {
	repo := repository.NewInMemoryRepo() //хранилище
	uc := NewTaskUsecase(repo)           //uc-то что мы проверяем(подключен или нет)

	err := uc.CreateTask("тест", "описание") //действие которое проверяем
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	tasks, _ := uc.ListTask() //достаем список задач
	if len(tasks) != 1 {      //ожидаем одну задачу
		t.Errorf("expected 1 task, got %d", len(tasks))
	}
}
