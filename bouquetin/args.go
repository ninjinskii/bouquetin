package bouquetin

import (
	"errors"
	"fmt"
	"strings"
)

const ARG_USER = "--user"
const ARG_TITLE = "--title"
const ARG_NOTE = "--note"

const COMMAND_ADD = "add"
const COMMAND_INIT = "init"

func ParseArgs(args []string) (string, string, string, string, string, error) {
	if len(args) == 0 {
		return "", "", "", "", "", errors.New("no args provided")
	}

	commands := []string{COMMAND_ADD, COMMAND_INIT}
	command := args[0]

	if !contains(commands, command) {
		return "", "", "", "", "", errors.New("wrong command")
	}

	var url, user, title, note string
	var skipNext = false

	for index, arg := range args {
		if skipNext {
			skipNext = false
			continue
		}

		if index == 0 {
			if !strings.HasPrefix(arg, "--") {
				url = arg
			}
		}

		// Avoid out of bound
		if len(args)-1 == index {
			continue
		}

		next := args[index+1]

		switch arg {
		case ARG_USER:
			if !strings.HasPrefix(next, "--") {
				skipNext = true
				user = args[index+1]
			}
		case ARG_TITLE:
			if !strings.HasPrefix(next, "--") {
				skipNext = true
				title = args[index+1]
			}
		case ARG_NOTE:
			if !strings.HasPrefix(next, "--") {
				skipNext = true
				note = args[index+1]
			}
		}
	}

	return command, url, user, title, note, nil
}

func Help() {
	fmt.Println("Usage:")
	fmt.Println("")
	fmt.Println("Add an entry with an already existing password")
	fmt.Println("bqt add [url] --user [username] --title [short description] --note [note]")
	fmt.Println("")
	fmt.Println("Add an entry, generate password and push it into your clipboard")
	fmt.Println("bqt new [url] --user [username] --title [short description] --note [note]")
	fmt.Println("")
	fmt.Println("Show all entries (without password)")
	fmt.Println("bqt ls")
	fmt.Println("")
	fmt.Println("Get a password, for the closest matched entry")
	fmt.Println("bqt [approx url or title]")
	fmt.Println("")
	fmt.Println("Remove an entry")
	fmt.Println("bqt rm [url]")
	fmt.Println("")
	fmt.Println("Sets up a key to store your passwords")
	fmt.Println("bqt init")
}

func contains(s []string, e string) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}
