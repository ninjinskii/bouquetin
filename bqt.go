package main

import (
	"fmt"
	"ninjinski/bouquetin/bouquetin"
	"os"
	"time"

	"golang.design/x/clipboard"
)

func main() {
	// var fTitle, fNote string
	// flag.StringVar(&fTitle, "t", "", "Simple short description of associated website(s)")
	// flag.StringVar(&fNote, "n", "", "Verbose description of associated website(s)")
	// flag.Parse()

	// fmt.Println(fTitle, fNote)

	// MAIN PACKAGE WILL BE IN CHARGE TO COPY DATA TO CLIPBOAR FOR W AND LINUX
	// (clipbaord doesnt work with android)

	argsWithoutProg := os.Args[1:]
	command, url, user, title, note, err := bouquetin.ParseArgs(argsWithoutProg)

	if err != nil {
		bouquetin.Help()
		return
	}
}

func testClipboard() {
	err := clipboard.Init()

	if err != nil {
		panic(err)
	}

	clipboard.Write(clipboard.FmtText, []byte("Clipboard works!"))
	a := string(clipboard.Read(clipboard.FmtText))
	fmt.Println(a)

	time.Sleep(8 * time.Second)
	fmt.Println("Fin")
}
