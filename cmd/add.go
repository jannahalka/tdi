/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"github.com/jannahalka/tdi/todo"
	"github.com/spf13/cobra"
)

var priority int

func addRun(cmd *cobra.Command, args []string) {
	// items, err := todo.ReadItems(datafile)

	items := []todo.Item{}

	for _, x := range args {
		item := todo.Item{Text:x}
		item.SetPriority(priority)
		items = append(items, item)
	}

	fmt.Println("Hello")

	err := todo.SaveItems(datafile, items)
	fmt.Println("Hello2")

	if err != nil {
		fmt.Errorf("%v", err)
	}
}

// addCmd represents the add command
var addCmd = &cobra.Command{
	Use:   "add",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: addRun,
}

func init() {
	rootCmd.AddCommand(addCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// addCmd.PersistentFlags().String("foo", "", "A help for foo")

	addCmd.Flags().IntVarP(&priority, "priority", "p", 2, "Priority:1,2,3")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// addCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
