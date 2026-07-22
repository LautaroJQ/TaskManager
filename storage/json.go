package storage

import (
	"TaskManager/models"
	"encoding/json"
	"os"
)

const defaultStorage string = "storage.json"

func SaveTask(task models.Task) error {
	binary, err := json.MarshalIndent(task, "", "	")

	if err != nil {
		return err
	}

	err = os.WriteFile(defaultStorage, binary, 0644)

	return err
}
