package models

import (
	"errors"
	"time"
)

// For creating a new task you need Title(Mandatory), Description(Optional), Duration(optional)

type Task struct {
	Id          int       `json:"id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	CreatedAt   time.Time `json:"created_at"`
	EndAt       time.Time `json:"end_at"`
}

const DefaultDuration int = 7

func CreateTask(Title string, Description string, Duration int) (Task, error) {
	if Title == "" {
		return Task{}, errors.New("Title cannot be empty")
	}

	if Duration <= 0 {
		Duration = DefaultDuration
	}

	Days := time.Hour * 24 * time.Duration(Duration)

	CreatedAt := time.Now()
	EndAt := CreatedAt.Add(Days)

	task := Task{
		0,
		Title,
		Description,
		CreatedAt,
		EndAt,
	}

	return task, nil
}
