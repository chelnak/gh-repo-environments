/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"log"

	"github.com/chelnak/gh-environments/internal/client"
	"github.com/chelnak/gh-environments/internal/cmd/delete"
	"github.com/spf13/cobra"
)

// deleteCmd represents the delete command
var deleteCmd = &cobra.Command{
	Use:   "delete <environment>",
	Short: "Delete an environment.",
	Long:  "Delete an environment.",
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {

		githubClient, err := client.NewClient()
		if err != nil {
			log.Fatal(err)
		}

		deletService := delete.NewDeleteService(githubClient)
		deleteOpts := delete.DeleteOptions{
			Name: args[0],
		}
		deletService.Delete(&deleteOpts)
	},
}

func init() {
	rootCmd.AddCommand(deleteCmd)
}
