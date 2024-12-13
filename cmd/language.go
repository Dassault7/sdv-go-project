/*
Copyright © 2024 JAROD GUICHARD
*/
package cmd

import (
	"github.com/spf13/cobra"
)

// languageCmd represents the language command
var languageCmd = &cobra.Command{
	Use:     "language",
	Aliases: []string{"lang", "l"},
	Short:   "List of languages",
	Long:    `List of languages provided by the API.`,
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Println("List of languages:")
		cmd.Println("\t1. en (English)")
		cmd.Println("\t2. fr (French)")
		cmd.Println("\t3. de (German)")
		cmd.Println("\t4. es (Spanish)")
		cmd.Println("\t5. cs (Czech)")
		cmd.Println("\t6. pt (Portuguese)")
	},
}

func init() {
	listCmd.AddCommand(languageCmd)
}
