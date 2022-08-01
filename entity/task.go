package entity

import "time"

type taskID int64
type TaskStatus string

const (
	TaskStatusTodo	TaskStatus = "todo"
	TaskStatusDoing	TaskStatus = "doing"
	TaskStatusDone	TaskStatus = "done"
)

type Task struct {
	ID			taskID			`json:"id"`
	Title		string			`json:"title"`
	Status	TaskStatus	`json:"status"`
	Created	time.Time		`json:"created"`
}

type Tasks []*Task
