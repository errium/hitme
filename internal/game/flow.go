package game

type Flow struct {
	state  uint8
	result uint8
}

const (
	StateNewGame uint8 = iota
	StatePlayerTurn
	StateDealerTurn
	StateGameOver
)

const (
	ResultPending uint8 = iota
	ResultPlayerWin
	ResultDealerWin
	ResultPush
)

func NewFlow() *Flow {
	return &Flow{
		state:  StateNewGame,
		result: ResultPending,
	}
}

func (f *Flow) State() uint8     { return f.state }
func (f *Flow) Result() uint8    { return f.result }
func (f *Flow) IsGameOver() bool { return f.state == StateGameOver }

func (f *Flow) Deal() {
	f.state = StatePlayerTurn
	f.result = ResultPending
}

func (f *Flow) Stand() {
	if f.state == StatePlayerTurn {
		f.state = StateDealerTurn
	}
}

func (f *Flow) EndGame(result uint8) {
	f.state = StateGameOver
	f.result = result
}

func (f *Flow) Reset() {
	f.state = StateNewGame
	f.result = ResultPending
}
