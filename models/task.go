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

func getNextID(tasks []Task) int {
	maxID := 0
	for _, t := range tasks {
		if t.Id > maxID {
			maxID = t.Id
		}
	}
	return maxID + 1
}

func CreateTask(Title string, Description string, Duration int, existingTasks []Task) (Task, error) {
	if Title == "" {
		return Task{}, errors.New("Title cannot be empty")
	}

	if Duration <= 0 {
		Duration = DefaultDuration
	}

	CreatedAt := time.Now()
	EndAt := CalculateEndDate(CreatedAt, Duration)
	Id := getNextID(existingTasks)
	task := Task{
		Id:          Id,
		Title:       Title,
		Description: Description,
		CreatedAt:   CreatedAt,
		EndAt:       EndAt,
		Done:        false,
	}

	return task, nil
}

func PrintTasks(tasks []Task) error {
	tabla := tablewriter.NewWriter(os.Stdout)
	tabla.Header([]string{"ID", "TITLE", "DESCRIPTION", "EXPIRATION DATE", "STATUS"})
	for _, task := range tasks {
		var status string = "Not completed"
		if CheckExpiration(task.EndAt) {
			status = "Expirated"
		}
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

func CheckExpiration(endDate time.Time) bool {
	if time.Now().Compare(endDate) == 1 {
		return true
	}
	return false
}
