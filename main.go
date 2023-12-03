package main

import (
	"fmt"
	"gol/cmd"
)

var (
	version = "dev"
	commit  = "none"
	date    = "unknown"
)

func main() {
	v := fmt.Sprintf("%s, commit %s, built at %s", version, commit, date)
	cmd.Execute(v)
}
