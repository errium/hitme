package game

type GameState int

const (
	NewGame GameState = iota
	PlayerTurn
	DealerTurn
	GameOver
)
