/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"os"
	"strconv"
	"text/tabwriter"

	"github.com/makiwebdeveloper/cli-todo/helpers"
	"github.com/spf13/cobra"
)

var completeCmd = &cobra.Command{
	Use:   "complete",
	Short: "Complete todo task -> cli-todo list [task id]",
	Run: func(cmd *cobra.Command, args []string) {
		tasks, err := helpers.LoadTasks("todo.csv")
		if err != nil {
			fmt.Println("Error:", err)
			return
		}

		id := args[0]

		existTask := 0

		for i := range tasks {
			if strconv.Itoa(tasks[i].ID) == id {
				existTask = 1

				tasks[i].Completed = true
				saveTasksError := helpers.SaveTasks("todo.csv", tasks)
				if saveTasksError != nil {
					fmt.Println("Error: ", saveTasksError)
					return
				}

				w := tabwriter.NewWriter(os.Stdout, 0, 0, 2, ' ', 0)

				fmt.Fprintln(w, "ID\tText\tCompleted")
				for _, task := range tasks {
					fmt.Fprintf(w, "%d\t%s\t%t\n", task.ID, task.Text, task.Completed)
				}

				w.Flush()
			}
		}
		if existTask == 0 {
			fmt.Println("Task with current ID doesnt exist")
		}
	},
}

func init() {
	rootCmd.AddCommand(completeCmd)
}
