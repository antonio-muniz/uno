package game

import (
	"github.com/antonio-muniz/uno/pkg/card"
	"github.com/antonio-muniz/uno/pkg/card/action"
	"github.com/antonio-muniz/uno/pkg/ui"
)

type Game struct {
	players *playerIterator
	deck    *deck
	pile    *pile
}

func New(players []Player) *Game {
	return &Game{
		players: newPlayerIterator(players),
		deck:    createUnoDeck(),
		pile:    createEmptyPile(),
	}
}

func (g *Game) Play() Player {
	g.dealStartingCards()
	g.playFirstCard()

	for {
		player := g.players.Next()
		gameState := g.extractGameState(player)
		card := player.Play(gameState, g.deck)
		if card != nil {
			g.pile.Add(card)
			g.performCardActions(card)
		}
		if player.NoCards() {
			return player.player
		}
	}
}

func (g *Game) dealStartingCards() {
	g.players.ForEach(func(player *playerController) {
		hand := g.deck.Draw(7)
		player.AddCards(hand)
	})
}

func (g *Game) playFirstCard() {
	firstCard := g.deck.DrawOne()
	g.pile.Add(firstCard)
	ui.Message.FirstCardPlayed(firstCard)
	g.performCardActions(firstCard)
}

func (g *Game) performCardActions(playedCard card.Card) {
	player := g.players.Current()
	for _, cardAction := range playedCard.Actions() {
		switch cardAction := cardAction.(type) {
		case action.DrawCardsAction:
			cards := g.deck.Draw(cardAction.Amount())
			g.players.Current().AddCards(cards)
		case action.ReverseTurnsAction:
			g.players.Reverse()
		case action.SkipTurnAction:
			g.players.Skip()
		case action.PickColorAction:
			gameState := g.extractGameState(player)
			color := player.PickColor(gameState)
			coloredCard := card.NewColoredCard(playedCard, color)
			g.pile.ReplaceTop(coloredCard)
			ui.Message.PlayerPickedColor(player.Name(), color)
		}
	}
}

func (g Game) extractGameState(player *playerController) GameState {
	playerSequence := make([]string, 0)
	playerHandCounts := make(map[string]int)

	g.players.ForEach(func(player *playerController) {
		playerSequence = append(playerSequence, player.Name())
		playerHandCounts[player.Name()] = len(player.Hand())
	})

	return GameState{
		lastPlayedCard:    g.pile.Top(),
		currentPlayerHand: player.Hand(),
		playerSequence:    playerSequence,
		playerHandCounts:  playerHandCounts,
	}
}
