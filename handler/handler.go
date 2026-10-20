package handler

import (
	"encoding/json"
	"net/http"
	"strconv"
	"todo-api/usecase"
)

type TaskHandler struct {
	uc usecase.TaskUsecase // интерфейс usecase-слоя, а не конкретная структура
}

func NewTaskHandler(uc1 usecase.TaskUsecase) *TaskHandler { // функция-конструктор, которая принимает интерфейс usecase-слоя и возвращает указатель на структуру TaskHandler
	return &TaskHandler{uc: uc1}
}

type createTaskRequest struct {
	Title       string `json:"title"`
	Description string `json:"description"`
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

func (h *TaskHandler) CreateTask(w http.ResponseWriter, r *http.Request) {
	var req createTaskRequest
	err := json.NewDecoder(r.Body).Decode(&req) // декодируем JSON из тела запроса в структуру createTaskRequest
	if err != nil {
		http.Error(w, "bad request", http.StatusBadRequest)
		return
	}

	task, err := h.uc.CreateTask(req.Title, req.Description) // передаём в usecase данные, которые Decoder положил в req
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)

	json.NewEncoder(w).Encode(task)

}
func (h *TaskHandler) DeleteTask(w http.ResponseWriter, r *http.Request) {
	idstr := r.PathValue("id")
	id, err := strconv.Atoi(idstr)
	if err != nil {
		http.Error(w, "invalid id", http.StatusBadRequest)
		return
	}
	err = h.uc.DeleteTask(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}

func (h *TaskHandler) GetTask(w http.ResponseWriter, r *http.Request) {
	idstr := r.PathValue("id")     // достаем значение сегмента айди из адреса
	id, err := strconv.Atoi(idstr) //перевод строки в число
	if err != nil {
		http.Error(w, "invalid id", http.StatusBadRequest) //task/abc-ошибка
		return
	}
	task, err := h.uc.GetTask(id) //спрашиваем задачу у юскасе
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(task)
}

func (h *TaskHandler) DoneTask(w http.ResponseWriter, r *http.Request) {
	idstr := r.PathValue("id")
	id, err := strconv.Atoi(idstr)
	if err != nil {
		http.Error(w, "invalid id", http.StatusBadRequest)
		return
	}
	task, err := h.uc.DoneTask(id) //спрашиваем задачу в юскасе
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(task)

}
