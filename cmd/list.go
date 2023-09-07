/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"os"

	"github.com/olistrik/temple/cmd/internal"
	"github.com/spf13/cobra"
)

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "list all available templates",
	Long:  "list all available templates",
	Args:  cobra.NoArgs,
	Run: func(cmd *cobra.Command, args []string) {
		templates, err := internal.LoadTemplates(internal.TEMPLATES, func(tmpl, path string) {})
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		fmt.Print("Templates:\n")
		for key := range templates {
			fmt.Printf("  %s\n", key)
		}
	},
}

func init() {
	rootCmd.AddCommand(listCmd)
}
