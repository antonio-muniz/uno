package game

import (
	"github.com/antonio-muniz/uno/game/card"
	"github.com/antonio-muniz/uno/game/card/color"
	"github.com/antonio-muniz/uno/game/ui"
)

type playerController struct {
	player Player
	hand   []card.Card
}

func newPlayerController(player Player) *playerController {
	return &playerController{
		player: player,
		hand:   make([]card.Card, 0, 7),
	}
}

func (c *playerController) AddCard(kard card.Card) {
	c.hand = append(c.hand, kard)
	c.player.NotifyCardsDrawn([]card.Card{kard})
}

func (c *playerController) AddCards(cards []card.Card) {
	c.hand = append(c.hand, cards...)
	c.player.NotifyCardsDrawn(cards)
}

func (c *playerController) Hand() []card.Card {
	hand := make([]card.Card, len(c.hand))
	copy(hand, c.hand)
	return hand
}

func (c *playerController) Name() string {
	return c.player.Name()
}

func (c *playerController) NoCards() bool {
	return len(c.hand) == 0
}

func (c *playerController) PickColor(gameState GameState) color.Color {
	return c.player.PickColor(gameState)
}

func (c *playerController) Play(gameState GameState) card.Card {
	playableCards := c.selectPlayableCards(gameState)
	if len(playableCards) == 0 {
		c.player.NotifyNoMatchingCardsInHand(gameState.LastPlayedCard(), gameState.CurrentPlayerHand())
		return nil
	}
	selectedCardIndex := c.player.Play(playableCards, gameState)
	selectedCard := c.hand[selectedCardIndex]
	c.hand[selectedCardIndex] = c.hand[len(c.hand)-1]
	c.hand = c.hand[:len(c.hand)-1]
	ui.Message.PlayerPlayedCard(c.Name(), selectedCard)
	return selectedCard
}

func (c *playerController) selectPlayableCards(gameState GameState) map[int]card.Card {
	currentColor := gameState.CurrentColor()
	lastPlayedCard := gameState.LastPlayedCard()

	playableCards := make(map[int]card.Card)
	for index, candidateCard := range c.hand {
		if Playable(candidateCard, currentColor, lastPlayedCard) {
			playableCards[index] = candidateCard
		}
	}
	return playableCards
}
