/*
Copyright © 2024 JAROD GUICHARD
*/
package cmd

import (
	"github.com/spf13/cobra"
)

// typesCmd represents the types command
var typesCmd = &cobra.Command{
	Use:   "types",
	Short: "List of types",
	Long:  `List of types provided by the API.`,
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Println("List of types:")
		cmd.Println("\t1. single")
		cmd.Println("\t2. twopart")
	},
}

func init() {
	listCmd.AddCommand(typesCmd)
}
