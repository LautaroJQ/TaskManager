/*
Copyright © 2026 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"TaskManager/storage"
	"fmt"
	"strconv"

	"github.com/spf13/cobra"
)

var (
	editTitle       string
	editDescription string
	editDuration    int
)

var editCmd = &cobra.Command{
	Use:   "edit <id>",
	Short: "This command let you edit an existing task",
	Long: `edit command let you edit an existing task
	e.g.:
	
		task edit 3 --title "New Title"
		task edit 5 --title "New Title" --description "New Description"
		
	With the correct ID you can change the title of an existing task`,

	Args: cobra.ExactArgs(1),

	RunE: func(cmd *cobra.Command, args []string) error {
		id, err := strconv.Atoi(args[0])

		if err != nil {
			fmt.Printf("An error has ocurred during id parsing: ", err)
			return err
		}

		err = storage.EditTask(id, editTitle, editDescription, editDuration, cmd)

		return err
	},
}

func init() {
	rootCmd.AddCommand(editCmd)

	editCmd.LocalFlags().StringVar(&editTitle, "title", "", "Indicates the new title of the task.")
	editCmd.LocalFlags().StringVar(&editDescription, "description", "", "Indicates the new description of the task.")
	editCmd.LocalFlags().IntVar(&editDuration, "duration", 0, "Indicates the new duration value")
}
