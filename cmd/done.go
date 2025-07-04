package cmd

import (
	"github.com/jannahalka/tdi/todo"
	"github.com/spf13/cobra"
	"log"
	"slices"
	"strconv"
)

func doneRun(cmd *cobra.Command, args []string) {
	items, err := todo.ReadItems(datafile)
	if err != nil {
		log.Fatalln(err)
	}

	for _, id := range args {
		id, err := strconv.Atoi(id)

		if err != nil {
			log.Fatalln(err)
		}

		idx := slices.IndexFunc(items, func(item todo.Item) bool {
			return item.Id == id

		})
		if idx == -1 {
			log.Fatalln("Could not find the id")
		}

		items[idx].SetToDone()
	}

	err = todo.SaveItems(datafile, items)
	if err != nil {
		log.Fatalln(err)
	}
}

var doneCmd = &cobra.Command{
	Use:   "done",
	Short: "Toggles a todo to completed.",
	Long:  "Command, which is used for toggling completed state to true.",
	Run:   doneRun,
}

func init() {
	rootCmd.AddCommand(doneCmd)
}
