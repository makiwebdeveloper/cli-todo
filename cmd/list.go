/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"os"
	"text/tabwriter"

	"github.com/makiwebdeveloper/cli-todo/helpers"
	"github.com/spf13/cobra"
)

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "Show todo list",
	Run: func(cmd *cobra.Command, args []string) {
		tasks, err := helpers.LoadTasks("todo.csv")
		if err != nil {
			fmt.Println("Error:", err)
			return
		}

		w := tabwriter.NewWriter(os.Stdout, 0, 0, 2, ' ', 0)

		fmt.Fprintln(w, "ID\tText\tCompleted")
		for _, task := range tasks {
			fmt.Fprintf(w, "%d\t%s\t%t\n", task.ID, task.Text, task.Completed)
		}

		w.Flush()
	},
}

func init() {
	rootCmd.AddCommand(listCmd)
}
