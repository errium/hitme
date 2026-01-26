package game

type Hand struct {
	Cards []Card
}

func NewHand() *Hand {
	return &Hand{Cards: make([]Card, 0)}
}
