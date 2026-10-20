package repository

import (
	"fmt"
	"sort"
)

type Task struct {
	ID          int
	Title       string
	Description string
	Status      string
}

type TaskRepository interface {
	Create(task Task) (Task, error) // принимает задачу и возвращает ошибку если что-то не так
	GetByID(id int) (Task, error)   // принимает айди и возвращает найденную задачу и ошибку
	List() ([]Task, error)          //принимает ничего и возвращает список задач
	Delete(id int) error
	Done(id int) (Task, error)
}

type InMemoryRepo struct {
	data   map[int]Task
	nextID int
}

func NewInMemoryRepo() *InMemoryRepo {
	return &InMemoryRepo{
		data:   make(map[int]Task),
		nextID: 1,
	}

}

func (r *InMemoryRepo) Delete(id int) error {
	_, ok := r.data[id]
	if !ok {
		return fmt.Errorf("task not found")
	} else {
		delete(r.data, id)
		return nil

	}

}
func (r *InMemoryRepo) Create(task Task) (Task, error) {
	task.ID = r.nextID
	r.nextID = r.nextID + 1
	r.data[task.ID] = task //Правильно так: «в map r.data по ключу task.ID кладём значение task». Ключ это ID, значение это сама задача.
	return task, nil
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

func (r *InMemoryRepo) List() ([]Task, error) {
	tasks := make([]Task, 0, len(r.data)) //создаём слайс задач, длина 0, ёмкость равна количеству задач в map
	for _, task := range r.data {         // ключ не нужен, поэтому используем _, task
		tasks = append(tasks, task) //добавляем каждую задачу в слайс
	}
	sort.Slice(tasks, func(i, j int) bool { //сортиуем слайс такс слайс меняется на месте присваивать результат не нужно
		return tasks[i].ID < tasks[j].ID //tasks[i].ID < tasks[j].ID — «элемент i идёт раньше j, если его ID меньше». Это сортировка по возрастанию ID.
	})
	return tasks, nil //возврашаем слайс задач и nil, если всё ок
}
func (r *InMemoryRepo) Done(id int) (Task, error) {
	task, ok := r.data[id] //ищем задачу по айди
	if !ok {
		return Task{}, fmt.Errorf("task not found")
	}

	task.Status = "done"   //меняем статус
	r.data[task.ID] = task //закидываем в мапу измененную задачу тобы изменение сохранилось, т.к. task это копия
	return task, nil
}
