package main

import (
	"fmt"
	"github.com/kulapard/gol/cmd"
)

var (
	version = "dev"
	commit  = "none"
	date    = "unknown"
)

func main() {
	version := fmt.Sprintf("%s, commit %s, built at %s", version, commit, date)
	cmd.Execute(version)
}
