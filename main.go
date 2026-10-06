package main

import (
	"net/http"
	"todo-api/handler"
	"todo-api/repository"
	"todo-api/usecase"
)

func main() {
	repo := repository.NewInMemoryRepo() //сoздаем репозиторий
	u := usecase.NewTaskUsecase(repo)    //usecase получает хранилище

	h := handler.NewTaskHandler(u) //handler получает usecase

	u.CreateTask("составить резюме", "срочно!")
	http.HandleFunc("/tasks", h.ListTasks)
	http.ListenAndServe(":8080", nil)

}
