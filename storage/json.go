package storage

import (
	"TaskManager/models"
	"encoding/json"
	"errors"
	"os"
)

const defaultStorageName string = "storage.json"

func SaveTasks(tasks []models.Task) error {
	binary, err := json.MarshalIndent(tasks, "", "	")

	if err != nil {
		return err
	}

	err = os.WriteFile(defaultStorageName, binary, 0644)

	return err
}

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

	err = SaveTasks(tasks)

	return err
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

func EditTask(id int, taskUpdate models.TaskUpdate) error {
	found := false
	tasks, err := GetTaskList()

	for index := range tasks {
		if tasks[index].Id != id {
			continue
		}

		found = true
		if taskUpdate.Title != nil {
			tasks[index].Title = *taskUpdate.Title
		}

		if taskUpdate.Description != nil {
			tasks[index].Description = *taskUpdate.Description
		}

		if taskUpdate.Duration != nil {
			tasks[index].EndAt = models.CalculateEndDate(tasks[index].CreatedAt, *taskUpdate.Duration)
		}
		break
	}

	if !found {
		return errors.New("ID not found")
	}

	err = SaveTasks(tasks)

	return err
}
