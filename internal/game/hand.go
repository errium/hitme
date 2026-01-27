package game

type Hand struct {
	Cards []Card
}

func NewHand() *Hand {
	return &Hand{Cards: make([]Card, 0)}
}

func (h *Hand) AddCard(c Card) {
	h.Cards = append(h.Cards, c)
}

func (h *Hand) Clear() {
	h.Cards = h.Cards[:0]
}

func (h *Hand) HandValue() int {
	total := 0
	aces := 0

	for _, c := range h.Cards {
		if c.Rank == "a" {
			aces++
			total += 11
		} else {
			total += c.CardValue()
		}
	}

	for total > 21 && aces > 0 {
		total -= 10
		aces--
	}

	return total
}

func (h *Hand) IsBlackjack() bool {
	return len(h.Cards) == 2 && h.HandValue() == 21
}

func (h *Hand) IsBusted() bool {
	return h.HandValue() > 21
}
