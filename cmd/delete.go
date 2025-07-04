/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"github.com/jannahalka/tdi/todo"
	"github.com/spf13/cobra"
	"log"
	"slices"
	"strconv"
)

func deleteRun(cmd *cobra.Command, args []string) {
	items, err := todo.ReadItems(datafile)

	if err != nil {
		log.Fatalln(err)
	}

	for _, id := range args {
		id, err := strconv.Atoi(id)
		if err != nil {
			log.Fatalln(err)
		}

		items = slices.DeleteFunc(items, func(item todo.Item) bool {
			return item.Id == id
		})
	}

	todo.SaveItems(datafile, items)
}

var deleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "Delete a todo.",
	Long:  "Deletes a todo using an id",
	Run:   deleteRun,
}

func init() {
	rootCmd.AddCommand(deleteCmd)
}
