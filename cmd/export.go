/*
Copyright © 2026 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"TaskManager/storage"
	"fmt"

	"github.com/spf13/cobra"
)

var (
	exportFileName string
)

// exportCmd represents the export command
var exportCmd = &cobra.Command{
	Use:   "export",
	Short: "Export to csv format",
	Long:  `This command let you export task list in csv format.`,
	Run: func(cmd *cobra.Command, args []string) {
		tasks, err := storage.GetTaskList()

		if err != nil {
			fmt.Println(err)
			return
		}

		storage.ExportCSV(tasks, exportFileName)
	},
}

func init() {
	rootCmd.AddCommand(exportCmd)

	exportCmd.LocalFlags().StringVar(&exportFileName, "export", "export", "set name of export file")
	exportCmd.MarkFlagRequired("export")
}
