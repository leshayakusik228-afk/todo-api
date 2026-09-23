package repository

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
