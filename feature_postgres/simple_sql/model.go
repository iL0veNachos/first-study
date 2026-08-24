package simple_sql

import (
	"fmt"
	"time"
)

type TaskModel struct {
	title        string
	description  string
	completed    bool
	created_at   time.Time
	completed_at *time.Time
}

func (t *TaskModel) Print() {
	fmt.Println("--------------------------")
	fmt.Println("title: ", t.title)
	fmt.Println("description: ", t.description)
	fmt.Println("completed: ", t.completed)
	fmt.Println("created_at: ", t.created_at)
	fmt.Println("completed_at: ", t.completed_at)
	fmt.Println("--------------------------")
}
