package repository

import "fmt"

type Task struct {
	ID          int
	Title       string
	Description string
	Status      string
}

type TaskRepository interface {
	Create(task Task) error       // принимает задачу и возвращает ошибку если что-то не так
	GetByID(id int) (Task, error) // принимает айди и возвращает найденную задачу и ошибку
	List() ([]Task, error)        //принимает ничего и возвращает список задач
}

type InMemoryRepo struct {
	data map[int]Task
}

func NewInMemoryRepo() *InMemoryRepo {
	return &InMemoryRepo{
		data: make(map[int]Task),
	}
}
func (r *InMemoryRepo) Create(task Task) error {
	r.data[task.ID] = task
	return nil
}

//(r *InMemoryRepo) — ресивер, указатель, чтобы менять реальную map r.data, а не копию
//Create(task Task) error — сигнатура точно совпадает с той, что в интерфейсе
//r.data[task.ID] = task — кладём переданную задачу в map по ключу task.ID
//return nil — сообщаем, что ошибки не было (успешно создали)

func (r *InMemoryRepo) GetByID(id int) (Task, error) {
	task, ok := r.data[id] //ищем задачу по айди в map, ok будет true если нашли, false если нет
	if !ok {
		return Task{}, fmt.Errorf("task not found") //возврашаем пустую задачу и ошибку

	}
	return task, nil //возврашаем найденную задачу и nil, если всё ок
}

func (r*InMemoryRepo) List() ([]Task, error) {
	tasks:=make([]Task,0,len(r.data))//создаём слайс задач, длина 0, ёмкость равна количеству задач в map
	for_,task:=range r.data{
		tasks=append(tasks,task)//добавляем каждую задачу в слайс
	}
	return tasks,nil//возврашаем слайс задач и nil, если всё ок	
	}