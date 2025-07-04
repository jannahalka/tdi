package cmd

import (
	"fmt"
	"github.com/jannahalka/tdi/todo"
	"github.com/spf13/cobra"
	"sort"
)

func listRun(cmd *cobra.Command, args []string) {
	items, err := todo.ReadItems(datafile)
	if err != nil {
		fmt.Println(err)
	}
	sort.Sort(todo.ByPri(items))
	todo.DisplayTodos(items)
}

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List all todos",
	Long:  "List all todos sorted based on priority",
	Run:   listRun,
}

func init() {
	rootCmd.AddCommand(listCmd)
}
