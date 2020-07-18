package game

import (
	"github.com/antonio-muniz/uno/pkg/card"
	"github.com/antonio-muniz/uno/pkg/card/color"
	"github.com/antonio-muniz/uno/pkg/ui"
)

type playerController struct {
	player Player
	hand   *Hand
}

func newPlayerController(player Player) *playerController {
	return &playerController{
		player: player,
		hand:   NewHand(),
	}
}

func (c *playerController) AddCards(cards []card.Card) {
	c.hand.AddCards(cards)
	c.player.NotifyCardsDrawn(cards)
}

func (c *playerController) Hand() []card.Card {
	return c.hand.Cards()
}

func (c *playerController) Name() string {
	return c.player.Name()
}

func (c *playerController) NoCards() bool {
	return c.hand.Empty()
}

func (c *playerController) PickColor(gameState GameState) color.Color {
	return c.player.PickColor(gameState)
}

func (c *playerController) Play(gameState GameState, deck *deck) card.Card {
	playableCards := c.hand.PlayableCards(gameState.LastPlayedCard())
	if len(playableCards) == 0 {
		c.player.NotifyNoMatchingCardsInHand(gameState.LastPlayedCard(), gameState.CurrentPlayerHand())
		return c.tryTopDecking(gameState, deck)
	}

	for {
		selectedCard := c.player.Play(playableCards, gameState)
		removed := c.hand.RemoveCard(selectedCard)
		if !removed {
			ui.Printfln("Cheat detected! Card %s is not in %s's hand!", selectedCard, c.player.Name())
			continue
		}
		ui.Message.PlayerPlayedCard(c.Name(), selectedCard)
		return selectedCard
	}
}

func (c *playerController) tryTopDecking(gameState GameState, deck *deck) card.Card {
	extraCard := deck.DrawOne()
	if Playable(extraCard, gameState.LastPlayedCard()) {
		ui.Message.PlayerDrewAndPlayedCard(c.Name(), extraCard)
		return extraCard
	}
	c.AddCards([]card.Card{extraCard})
	ui.Message.PlayerPassed(c.Name())
	return nil
}
