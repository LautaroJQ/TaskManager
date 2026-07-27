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

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "Shows you a list of tasks an its status",
	Long:  `Shows you a list of tasks an its status`,
	Run: func(cmd *cobra.Command, args []string) {
		tasks, err := storage.GetTaskList()

		if err != nil {
			fmt.Println(err)
			return
		}

		models.PrintTasks(tasks)
	},
}

func init() {
	rootCmd.AddCommand(listCmd)
}
