package storage

import (
	"TaskManager/models"
	"encoding/json"
	"os"
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
