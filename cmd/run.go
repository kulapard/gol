package cmd

import (
	"github.com/spf13/cobra"
	"gol/game"
)

var speed int
var fileName string
var rows, cols int

var cmdRun = &cobra.Command{
	Use:   "run",
	Short: "Run Game of Life",
	Run: func(cmd *cobra.Command, args []string) {
		game.RunGameOfLife(fileName, speed, rows, cols)
	},
}

func init() {
	rootCmd.AddCommand(cmdRun)
	cmdRun.Flags().StringVarP(&fileName, "file", "f", "", "path to the file with initial state")
	cmdRun.Flags().IntVarP(&speed, "speed", "s", 5, "evolution speed in generations per second")
	cmdRun.Flags().IntVarP(&rows, "rows", "r", 40, "number of rows to generate (ignored if file is provided)")
	cmdRun.Flags().IntVarP(&cols, "cols", "c", 40, "number of cols to generate (ignored if file is provided)")
}
