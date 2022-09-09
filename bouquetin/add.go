package bouquetin

import (
	"fmt"
	"log"
)

type AddState struct {
	bqt *Bqt
}

func (state *AddState) run() {
	if checkInit() {
		log.Fatalln("Not implemented yet.")
	} else {
		_, err := state.bqt.retryWithState(state.bqt.initState)

		if err != nil {
			log.Fatalln("Not implemented yet.")
		}

		// Init gone wrong. Giving up.
		fmt.Println("An error occured while initilaizing database.")
	}
}
