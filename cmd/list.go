/*
Copyright © 2024 JAROD GUICHARD
*/
package cmd

import (
	"github.com/spf13/cobra"
)

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List of items",
	Long:  `List of items provided by the API, for different categories.`,
}

func init() {
	rootCmd.AddCommand(listCmd)
}
