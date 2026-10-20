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
	u.CreateTask("купить продукты", "молоко, хлеб, яйца")
	u.CreateTask("оплатить интернет", "до конца недели")
	u.CreateTask("записаться к стоматологу", "позвонить утром")
	u.CreateTask("постирать вещи", "белое и тёмное отдельно")
	u.CreateTask("позвонить родителям", "в воскресенье вечером")
	u.CreateTask("сходить в спортзал", "тренировка ног")
	u.CreateTask("заменить масло в машине", "записаться в сервис")
	http.HandleFunc("GET /tasks/{id}", h.GetTask)
	http.HandleFunc("DELETE /tasks/{id}", h.DeleteTask)
	http.HandleFunc("PATCH /tasks/{id}/done", h.DoneTask)
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
