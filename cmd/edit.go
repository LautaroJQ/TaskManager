/*
Copyright © 2026 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"TaskManager/models"
	"TaskManager/storage"
	"errors"
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

		if !cmd.Flags().Changed("title") &&
			!cmd.Flags().Changed("description") &&
			!cmd.Flags().Changed("duration") {
			return errors.New("no fields to update")
		}

		if editTitle == "" && cmd.Flags().Changed("title") {
			return errors.New("Could not have an empty title")
		}

		id, err := strconv.Atoi(args[0])

		if err != nil {
			fmt.Printf("An error has ocurred during id parsing: %v", err)
			return err
		}

		updates := models.TaskUpdate{}

		if cmd.Flags().Changed("title") {
			updates.Title = &editTitle
		}

		if cmd.Flags().Changed("description") {
			updates.Description = &editDescription
		}

		if cmd.Flags().Changed("duration") {
			updates.Duration = &editDuration
		}

		err = storage.EditTask(id, updates)

		return err
	},
}

func init() {
	rootCmd.AddCommand(editCmd)

	editCmd.Flags().StringVar(&editTitle, "title", "", "Indicates the new title of the task.")
	editCmd.Flags().StringVar(&editDescription, "description", "", "Indicates the new description of the task.")
	editCmd.Flags().IntVar(&editDuration, "duration", 0, "Indicates the new duration value")
}
