package models

type Task struct {
	ID        int    `csv:id`
	Text      string `csv:text`
	Completed bool   `csv:completed`
}
