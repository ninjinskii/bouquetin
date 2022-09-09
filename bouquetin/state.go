package bouquetin

type State interface {
	run() (string, error)
}

func (bqt *Bqt) setState(state State) *Bqt {
	bqt.currentState = state
	return bqt
}

func (bqt *Bqt) retryWithState(state State) (string, error) {
	return bqt.setState(state).currentState.run()
}
