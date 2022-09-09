package bouquetin

import (
	"fmt"
)

type InitState struct {
	bqt *Bqt
}

func (state *InitState) run() {
	if checkInit() {
		fmt.Println("You're all set. Init not needed")
	} else {
		// prompt pwd x 2
		// check pwd is same
		// setup db
		// newMasterKey = nil
	}
}

func checkInit() bool {
	return false
}
