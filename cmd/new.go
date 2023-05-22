/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"io/ioutil"
	"os"
	"path"

	"git.hub.klippa.com/klippa/temple/cmd/internal"
	"github.com/spf13/cobra"
)

// newCmd represents the new command
var newCmd = &cobra.Command{
	Use:     "new",
	Aliases: []string{"n"},
	Short:   "Create a new temple from a template",
	Long:    "Create a new temple from a template",
	Run: func(cmd *cobra.Command, args []string) {
		files := []internal.File{}

		templates, err := internal.LoadTemplates(internal.TEMPLATES, func(tmpl, path string) {
			files = append(files, internal.File{
				Template: tmpl,
				Path:     path,
			})
		})
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		tmpl, ok := templates[args[0]]
		if !ok {
			fmt.Printf("%s is not a template\n", args[0])
			os.Exit(1)
		}

		// fill datamap with args
		datamap := map[string]string{}
		for i, d := range *&tmpl.Arguments {
			datamap[d.Key] = args[1+i]
		}

		// run template.
		err = tmpl.Template.Execute(ioutil.Discard, datamap)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		for _, file := range files {
			os.MkdirAll(path.Dir(file.Path), os.ModePerm)
			writer, err := os.Create(file.Path)
			if err != nil {
				fmt.Println(err)
				os.Exit(1)
			}

			err = tmpl.Template.ExecuteTemplate(writer, file.Template, datamap)
			if err != nil {
				fmt.Println(err)
				os.Exit(1)
			}

			fmt.Printf("template %s written to %s\n", file.Template, file.Path)
		}
	},
}

func init() {
	rootCmd.AddCommand(newCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// newCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// newCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
