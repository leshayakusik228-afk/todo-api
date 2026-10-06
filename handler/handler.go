package handler

import (
	"encoding/json"
	"net/http"
	"todo-api/usecase"
)

type TaskHandler struct {
	uc usecase.TaskUsecase // интерфейс usecase-слоя, а не конкретная структура
}

func NewTaskHandler(uc1 usecase.TaskUsecase) *TaskHandler { // функция-конструктор, которая принимает интерфейс usecase-слоя и возвращает указатель на структуру TaskHandler
	return &TaskHandler{uc: uc1}
}
func (h *TaskHandler) ListTasks(w http.ResponseWriter, r *http.Request) {
	tasks, err := h.uc.ListTask() // спрашиваем у usecase список задач
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(tasks) // кодируем слайс задач в JSON и отправляем в ответ
}
