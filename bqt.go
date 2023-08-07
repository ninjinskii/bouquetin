package main

import (
	"ninjinski/bouquetin/cmd"
	"os"
)

func main() {
	// TODO: remove for production
	// os.Setenv("BOUQUETIN_ID", "1234-5678")
	//  village 2 minde de fer 12h42
	os.Setenv("BOUQUETIN_FILEPATH", "./README.md")
	cmd.Execute()
}
