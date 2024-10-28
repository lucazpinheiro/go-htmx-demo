package internal

import "fmt"

type Task struct {
	ID          int    `json:"id"`
	Description string `json:"description"`
	Done        bool   `json:"done"`
}

type TaskRepository struct {
	TaskList []Task
}

func NewTaskRepository() TaskRepository {
	return TaskRepository{
		TaskList: []Task{
			{
				ID:          0,
				Description: "A",
				Done:        false,
			},
			{
				ID:          1,
				Description: "B",
				Done:        false,
			},
			{
				ID:          2,
				Description: "C",
				Done:        false,
			},
		},
	}
}

func (t *TaskRepository) CreateTask(description string) {
	fmt.Println(t.TaskList)
	newTask := Task{
		ID:          t.TaskList[len(t.TaskList)-1].ID + 1,
		Description: description,
	}
	t.TaskList = append(t.TaskList, newTask)
}

func (t *TaskRepository) ListTask() []Task {
	return t.TaskList
}

func (t *TaskRepository) FlipTaskStatus(id int) {
	for i, task := range t.TaskList {
		if task.ID == id {
			t.TaskList[i].Done = !task.Done
			break
		}
	}
}
