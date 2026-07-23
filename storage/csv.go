package storage

import (
	"TaskManager/models"
	"encoding/csv"
	"os"
	"strconv"
)

func ExportCSV(tasks []models.Task, filename string) error {
	f, err := os.OpenFile(filename+".csv", os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}
	defer f.Close()

	w := csv.NewWriter(f)
	header := []string{"ID", "TITLE", "DESCRIPTION", "CREATION DATE", "END DATE", "DONE"}
	if err := w.Write(header); err != nil {
		return err
	}
	for _, task := range tasks {
		row := []string{
			strconv.Itoa(task.Id),
			task.Title,
			task.Description,
			task.CreatedAt.Format("02-01-2006"),
			task.EndAt.Format("02-01-2006"),
			strconv.FormatBool(task.Done),
		}
		err = w.Write(row)

		if err != nil {
			return err
		}
	}

	w.Flush()

	err = w.Error()

	return err
}
