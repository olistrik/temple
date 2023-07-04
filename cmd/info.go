/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"io/ioutil"
	"os"

	"github.com/kranex/temple/cmd/internal"
	"github.com/spf13/cobra"
)

// infoCmd represents the info command
var infoCmd = &cobra.Command{
	Use:   "info <template>",
	Short: "show information on a template",
	Long:  `show information on a template`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("info called")
		templates, err := internal.LoadTemplates(internal.TEMPLATES, func(tmpl, path string) {})
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		tmpl, ok := templates[args[0]]
		if !ok {
			fmt.Printf("unknown template %s\n", args[0])
			os.Exit(1)
		}

		// preload data from template `declare`
		err = tmpl.Template.Execute(ioutil.Discard, nil)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		fmt.Printf("%s\n\n", tmpl.Desciption)

		fmt.Printf("Usage:\n  temple new %s", args[0])

		for _, dat := range tmpl.Arguments {
			fmt.Printf(" [%s]", dat.Key)
		}

		fmt.Print("\n\nArguments:\n")

		for _, dat := range tmpl.Arguments {
			fmt.Printf("  %s\t\t%s\n", dat.Key, dat.Desciption)
		}

	},
}

func init() {
	rootCmd.AddCommand(infoCmd)
}
