package main

import (
	"gol/cmd"
)

var (
	version = "unknown-local-build"
)

func main() {
	cmd.Execute(version)
}
