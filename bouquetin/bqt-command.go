package bouquetin

import "log"

type Bqt struct {
	addState     State
	newState     State
	lsState      State
	rmState      State
	initState    State
	currentState State

	user         string
	title        string
	note         string
	query        string
	newMasterKey string // used when init
}

func makeCommand(user string, title string, note string, query string) *Bqt {
	bqt := &Bqt{
		user:  user,
		title: title,
		note:  note,
		query: query,
	}

	addState := &AddState{
		bqt: bqt,
	}

	initState := &InitState{
		bqt: bqt,
	}

	bqt.setState(nil)
	addState = addState
	initState = initState

	return bqt
}

func (bqt *Bqt) Run() (string, error) {
	if bqt.currentState != nil {
		return bqt.currentState.run()
	} else {
		log.Fatalln("Set a state to Bqt before trying to call Run().")
		return "", nil
	}
}
