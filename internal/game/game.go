package game

type Game struct {
	flow       *Flow
	deck       *Deck
	playerHand *Hand
	dealerHand *Hand
}

const dealerStandValue int = 17

func NewGame() *Game {
	return &Game{
		flow:       NewFlow(),
		deck:       NewDeck(),
		playerHand: NewHand(),
		dealerHand: NewHand(),
	}
}

func (g *Game) Deal() {
	g.deck.Shuffle()

	for range 2 {
		g.playerHand.AddCard(g.deck.Draw())
		g.dealerHand.AddCard(g.deck.Draw())
	}

	g.flow.Deal()
}

func (g *Game) Hit() {
	if g.flow.State() == StatePlayerTurn {
		g.playerHand.AddCard(g.deck.Draw())
		if g.playerHand.IsBusted() {
			g.flow.EndGame(ResultDealerWin)
		}
	}
}

func (g *Game) Stand() {
	if g.flow.State() != StatePlayerTurn {
		return
	}

	g.flow.Stand()
	g.DealerPlay()

	winner := g.DetermineWinner()

	g.flow.EndGame(winner)
}

func (g *Game) DealerPlay() {
	for g.dealerHand.Value() < dealerStandValue {
		g.dealerHand.AddCard(g.deck.Draw())
	}
}

func (g *Game) DetermineWinner() uint8 {
	playerScore := g.playerHand.Value()
	dealerScore := g.dealerHand.Value()

	if g.dealerHand.IsBusted() {
		return ResultPlayerWin
	}

	if g.playerHand.IsBusted() {
		return ResultDealerWin
	}

	if playerScore > dealerScore {
		return ResultPlayerWin
	} else if dealerScore > playerScore {
		return ResultDealerWin
	}

	return ResultTie
}
