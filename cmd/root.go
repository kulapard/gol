package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
)

var rootCmd = &cobra.Command{
	Use: "game",
	// TODO: write description
	Short:             "Game of life",
	Long:              `Game of life...`,
	CompletionOptions: cobra.CompletionOptions{DisableDefaultCmd: true},
}

func Execute(version string) {
	rootCmd.Version = version
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
