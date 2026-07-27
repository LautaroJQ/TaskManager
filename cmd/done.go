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

// doneCmd represents the done command
var doneCmd = &cobra.Command{
	Use:   "done <id>",
	Short: "Mark a task as done",
	Long:  `This command let you mark a task as done `,

	Args: cobra.ExactArgs(1),

	Run: func(cmd *cobra.Command, args []string) {
		id, err := strconv.Atoi(args[0])

		if err != nil {
			fmt.Printf("An error has ocurred during id parsing: %v", err)
			return
		}

		err = storage.MarkTaskDone(id)

		if err != nil {
			fmt.Printf("An error has ocurred: %v", err)
		}
	},
}

func init() {
	rootCmd.AddCommand(doneCmd)
}
