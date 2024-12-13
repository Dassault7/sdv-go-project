/*
Copyright © 2024 JAROD GUICHARD
*/
package cmd

import (
	"github.com/spf13/cobra"
)

// categoriesCmd represents the categories command
var categoriesCmd = &cobra.Command{
	Use:     "categories",
	Aliases: []string{"category", "cat", "c"},
	Short:   "List of categories",
	Long:    `List of categories provided by the API.`,
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Println("List of categories:")
		cmd.Println("\t1. any")
		cmd.Println("\t2. misc")
		cmd.Println("\t3. programming")
		cmd.Println("\t4. pun")
		cmd.Println("\t5. spooky")
		cmd.Println("\t6. christmas")
	},
}

func init() {
	listCmd.AddCommand(categoriesCmd)
}
