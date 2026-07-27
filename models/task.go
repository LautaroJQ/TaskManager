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
	Done        bool      `json:"done"`
}

const DefaultDuration int = 7

func CreateTask(Title string, Description string, Duration int) (Task, error) {
	if Title == "" {
		return Task{}, errors.New("Title cannot be empty")
	}

	if Duration <= 0 {
		Duration = DefaultDuration
	}

	CreatedAt := time.Now()
	EndAt := CalculateEndDate(CreatedAt, Duration)

	task := Task{
		0,
		Title,
		Description,
		CreatedAt,
		EndAt,
		false,
	}

	return task, nil
}

func PrintTasks(tasks []Task) error {
	tabla := tablewriter.NewWriter(os.Stdout)
	tabla.Header([]string{"ID", "TITLE", "DESCRIPTION", "EXPIRATION DATE", "STATUS"})
	for _, task := range tasks {
		var status string = "Not completed"
		if task.Done {
			status = "Completed"
		}
		err := tabla.Append(
			[]string{
				strconv.Itoa(task.Id),
				task.Title,
				task.Description,
				task.EndAt.Format("02-01-2006"),
				status,
			},
		)
		if err != nil {
			return err
		}
	}

	err := tabla.Render()
	return err
}

func CalculateEndDate(CreatedAt time.Time, duration int) time.Time {
	Days := time.Hour * 24 * time.Duration(duration)
	return CreatedAt.Add(Days)
}
