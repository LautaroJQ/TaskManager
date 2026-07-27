/*
Copyright © 2026 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"TaskManager/models"
	"TaskManager/storage"
	"fmt"

	"github.com/spf13/cobra"
)

var (
	Title       string
	Description string
	Duration    int
)

// addCmd represents the add command
var addCmd = &cobra.Command{
	Use:   "add",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		tasks, err := storage.GetTaskList()

		if err != nil {
			fmt.Print("An error has ocurred.")
		}

		task, err := models.CreateTask(Title, Description, Duration, tasks)

		if err != nil {
			fmt.Println(err)
			return
		}

		err = storage.SaveTask(task)

		if err != nil {
			fmt.Println(err)
			return
		}
	},
}

func init() {
	rootCmd.AddCommand(addCmd)

	addCmd.Flags().StringVar(&Title, "title", "", "Task title")
	addCmd.MarkFlagRequired("title")
	addCmd.Flags().StringVar(&Description, "description", "-", "Task description")
	addCmd.Flags().IntVar(&Duration, "duration", 0, "Task duration")
}
