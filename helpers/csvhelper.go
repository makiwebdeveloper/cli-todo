package helpers

import (
	"encoding/csv"
	"os"
	"strconv"

	"github.com/makiwebdeveloper/cli-todo/models"
)

func LoadTasks(filePath string) ([]models.Task, error) {
	var tasks []models.Task

	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		file, err := os.Create(filePath)
		if err != nil {
			return nil, err
		}
		defer file.Close()
		return tasks, nil
	}

	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	reader := csv.NewReader(file)
	records, err := reader.ReadAll()
	if err != nil {
		return nil, err
	}

	for _, record := range records {
		id, _ := strconv.Atoi(record[0])
		completed, _ := strconv.ParseBool(record[2])
		tasks = append(tasks, models.Task{
			ID:        id,
			Text:      record[1],
			Completed: completed,
		})
	}

	return tasks, nil
}

func SaveTasks(filePath string, tasks []models.Task) error {
	file, err := os.Create(filePath)
	if err != nil {
		return err
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	for _, task := range tasks {
		err := writer.Write([]string{
			strconv.Itoa(task.ID),
			task.Text,
			strconv.FormatBool(task.Completed),
		})
		if err != nil {
			return err
		}
	}

	return nil
}
