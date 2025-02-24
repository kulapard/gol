package main

import (
	"fmt"
	"github.com/kulapard/gol/core"
	"github.com/spf13/cobra"
	"os"
)

var (
	revision = "unknown"
)

var rootCmd = &cobra.Command{
	Use:               "gol",
	Short:             "Game of life",
	Long:              "Conway's Game of Life written in Go. \nSee https://en.wikipedia.org/wiki/Conway%27s_Game_of_Life for more info about the game.",
	CompletionOptions: cobra.CompletionOptions{DisableDefaultCmd: true},
}

var speed int
var fileName string
var rows, cols int

var cmdRun = &cobra.Command{
	Use:   "run",
	Short: "Run Game of Life",
	Run: func(_ *cobra.Command, _ []string) {
		gol, err := core.SetupGameOfLife(fileName, speed, rows, cols)
		if err != nil {
			fmt.Println(err)
			return
		}
		gol.RunForever()
	},
}

func main() {
	rootCmd.Version = revision

	rootCmd.AddCommand(cmdRun)
	cmdRun.Flags().StringVarP(&fileName, "file", "f", "", "path to the file with initial state")
	cmdRun.Flags().IntVarP(&speed, "speed", "s", 5, "evolution speed in generations per second")
	cmdRun.Flags().IntVarP(&rows, "rows", "r", 30, "number of rows to generate (ignored if file is provided)")
	cmdRun.Flags().IntVarP(&cols, "cols", "c", 40, "number of cols to generate (ignored if file is provided)")

	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
