package cmd

import (
	"github.com/jannahalka/tdi/todo"
	"github.com/spf13/cobra"
	"log"
)

var priority int

func addRun(cmd *cobra.Command, args []string) {
	items, err := todo.ReadItems(datafile)
	if err != nil {
		log.Fatal(err)
	}

	for _, x := range args {
		var lastId int
		n := len(items)

		if n == 0 {
			lastId = 0
		} else {
			lastId = items[n-1].Id
		}

		item := todo.Item{Text: x, Id: lastId + 1, Done: false}
		item.SetPriority(priority)
		items = append(items, item)
	}

	err = todo.SaveItems(datafile, items)

	if err != nil {
		log.Fatal("%v", err)
	}
}

var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Add a todo",
	Long:  "Adding a todo",
	Run:   addRun,
}

func init() {
	rootCmd.AddCommand(addCmd)

	addCmd.Flags().IntVarP(&priority, "priority", "p", 2, "Priority:1,2,3")
}
