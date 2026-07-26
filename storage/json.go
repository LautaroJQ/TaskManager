package storage

import (
	"TaskManager/models"
	"encoding/json"
	"errors"
	"os"
	"time"

	"github.com/spf13/cobra"
)

const defaultStorageName string = "storage.json"

func SaveTask(task models.Task) error {
	tasks := []models.Task{}
	content, err := os.ReadFile(defaultStorageName)

	if err != nil {
		if !os.IsNotExist(err) {
			return err
		}
	} else if len(content) > 0 {
		if err := json.Unmarshal(content, &tasks); err != nil {
			return err
		}
	}

	tasks = append(tasks, task)

	binary, err := json.MarshalIndent(tasks, "", "	")

	if err != nil {
		return err
	}

	os.WriteFile(defaultStorageName, binary, 0644)

	return nil
}

func GetTaskList() ([]models.Task, error) {
	tasks := []models.Task{}

	content, err := os.ReadFile(defaultStorageName)

	if err != nil {
		return []models.Task{}, err
	}

	err = json.Unmarshal(content, &tasks)

	if err != nil {
		return nil, err
	}

	return tasks, nil
}

func EditTask(id int, title string, description string, duration int, cmd *cobra.Command) error {

	content, err := os.ReadFile(defaultStorageName)

	if err != nil {
		return err
	}

	tasks := []models.Task{}
	task := models.Task{}
	err = json.Unmarshal(content, &tasks)

	if err != nil {
		return err
	}

	for _, task = range tasks {
		if task.Id == id {
			break
		}
	}

	if task.Id != id {
		return errors.New("There is no task with the input ID.")
	}

	if cmd.Flags().Changed("title") {
		task.Title = title
	}

	if cmd.Flags().Changed("description") {
		task.Description = description
	}

	if cmd.Flags().Changed("duration") {
		Days := time.Hour * 24 * time.Duration(duration)
		task.EndAt = task.CreatedAt.Add(Days)
	}

	return nil
}
