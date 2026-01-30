package game

type Game struct {
	Flow       *Flow
	deck       *Deck
	PlayerHand *Hand
	DealerHand *Hand
}

const dealerStandValue int = 17

func NewGame() *Game {
	return &Game{
		Flow:       NewFlow(),
		deck:       NewDeck(),
		PlayerHand: NewHand(),
		DealerHand: NewHand(),
	}
}

func (g *Game) Deal() {
	g.deck.Shuffle()

	for range 2 {
		g.PlayerHand.AddCard(g.deck.Draw())
		g.DealerHand.AddCard(g.deck.Draw())
	}

	g.Flow.Deal()
}

func (g *Game) Hit() {
	if g.Flow.State() == StatePlayerTurn {
		g.PlayerHand.AddCard(g.deck.Draw())
		if g.PlayerHand.IsBusted() {
			g.Flow.EndGame(ResultDealerWin)
		}
	}
}

func (g *Game) Stand() {
	if g.Flow.State() != StatePlayerTurn {
		return
	}

	g.Flow.Stand()
	g.DealerPlay()

	winner := g.DetermineWinner()

	g.Flow.EndGame(winner)
}

func (g *Game) DealerPlay() {
	for g.DealerHand.Value() < dealerStandValue {
		g.DealerHand.AddCard(g.deck.Draw())
	}
}

func (g *Game) DetermineWinner() uint8 {
	playerScore := g.PlayerHand.Value()
	dealerScore := g.DealerHand.Value()

	if g.DealerHand.IsBusted() {
		return ResultPlayerWin
	}

	if g.PlayerHand.IsBusted() {
		return ResultDealerWin
	}

	if playerScore > dealerScore {
		return ResultPlayerWin
	} else if dealerScore > playerScore {
		return ResultDealerWin
	}

	return ResultPush
}
