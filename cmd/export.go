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
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
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
