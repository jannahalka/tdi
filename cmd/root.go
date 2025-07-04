package cmd

import (
	"github.com/spf13/cobra"
	"log"
	"os"
)

var datafile string

var rootCmd = &cobra.Command{
	Use:   "tdi",
	Short: "Todo application",
	Long:  "Todo application created by following tutorial from: https://spf13.com/presentation/building-an-awesome-cli-app-in-go-oscon/",
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	home, err := os.UserHomeDir()

	if err != nil {
		log.Println("Unable to detect home directory. Please set data file using --datafile")
	}

	rootCmd.PersistentFlags().StringVar(&datafile, "datafile", home+string(os.PathSeparator)+".todos.csv", "datafile to store todos")
}
