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
	http.HandleFunc("GET /tasks/{id}", h.GetTask)
	http.HandleFunc("/tasks", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			h.ListTasks(w, r)
		case http.MethodPost:
			h.CreateTask(w, r)
		default:
			http.Error(w, "metod not allwed", http.StatusMethodNotAllowed)
		}
	})

	http.ListenAndServe(":8080", nil)

}
