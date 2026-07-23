package models

import (
	"errors"
	"os"
	"strconv"
	"time"

	"github.com/olekukonko/tablewriter"
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

func PrintTasks(tasks []Task) error {
	tabla := tablewriter.NewWriter(os.Stdout)
	tabla.Header([]string{"ID", "TITLE", "DESCRIPTION", "EXPIRATION DATE"})
	for i := range tasks {
		err := tabla.Append([]string{strconv.Itoa(tasks[i].Id), tasks[i].Title, tasks[i].Description, tasks[i].EndAt.Format("'02-01-2006")})
		if err != nil {
			return err
		}
	}

	err := tabla.Render() // Imprime la tabla directo en pantalla
	return err
}
