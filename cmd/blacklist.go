/*
Copyright © 2024 JAROD GUICHARD
*/
package cmd

import (
	"github.com/spf13/cobra"
)

// blacklistCmd represents the blacklist command
var blacklistCmd = &cobra.Command{
	Use:     "blacklist",
	Aliases: []string{"black", "bl", "b"},
	Short:   "List of blacklisted categories",
	Long:    `List of blacklisted categories provided by the API.`,
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Println("List of blacklisted categories:")
		cmd.Println("\t1. nsfw")
		cmd.Println("\t2. religious")
		cmd.Println("\t3. political")
		cmd.Println("\t4. racist")
		cmd.Println("\t5. sexist")
		cmd.Println("\t6. explicit")
	},
}

func init() {
	listCmd.AddCommand(blacklistCmd)
}
