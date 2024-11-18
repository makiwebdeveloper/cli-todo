/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/makiwebdeveloper/cli-todo/helpers"
	"github.com/makiwebdeveloper/cli-todo/models"
	"github.com/spf13/cobra"
)

var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Create new todo",
	Run: func(cmd *cobra.Command, args []string) {
		text := args[0]

		tasks, err := helpers.LoadTasks("todo.csv")
		if err != nil {
			fmt.Println("No tasks found")
			return
		}

		var newTask models.Task

		if len(tasks) > 0 {
			newTask = models.Task{
				ID:        tasks[len(tasks)-1].ID + 1,
				Text:      text,
				Completed: false,
			}
		} else {
			newTask = models.Task{
				ID:        1,
				Text:      text,
				Completed: false,
			}
		}

		tasks = append(tasks, newTask)

		saveTasksError := helpers.SaveTasks("todo.csv", tasks)
		if saveTasksError != nil {
			fmt.Println("Error: ", saveTasksError)
			return
		}

		fmt.Println("Added task:", newTask)
	},
}

func init() {
	rootCmd.AddCommand(addCmd)
}
