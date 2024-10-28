package internal

import (
	"reflect"
	"testing"
)

func TestTasksRepository(t *testing.T) {
	t.Run("create new task", func(t *testing.T) {
		want := []Task{
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
			{
				ID:          3,
				Description: "x",
				Done:        false,
			},
			{
				ID:          4,
				Description: "y",
				Done:        false,
			},
		}

		tasksRepo := NewTaskRepository()

		tasksRepo.CreateTask("x")
		tasksRepo.CreateTask("y")

		got := tasksRepo.ListTask()

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v, want %v", got, want)
		}
	})
}
