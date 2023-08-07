package main

import (
	"ninjinski/bouquetin/cmd"
	"os"
)

func main() {
	// TODO: remove for production
	os.Setenv("BOUQUETIN_FILEPATH", "./README.md")
	cmd.Execute()
}
