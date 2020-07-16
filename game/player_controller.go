package game

import (
	"errors"

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
	c.AddCards([]card.Card{kard})
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

func (c *playerController) Play(gameState GameState, deck *deck) card.Card {
	playableCards := c.selectPlayableCards(gameState)
	if len(playableCards) == 0 {
		c.player.NotifyNoMatchingCardsInHand(gameState.LastPlayedCard(), gameState.CurrentPlayerHand())
		return c.tryTopDecking(gameState, deck)
	}

	for {
		selectedCard := c.player.Play(playableCards, gameState)
		err := c.removeCardFromHand(selectedCard)
		if err != nil {
			ui.Printfln("Cheat detected! Card %s is not in %s's hand!", selectedCard, c.player.Name())
			continue
		}
		ui.Message.PlayerPlayedCard(c.Name(), selectedCard)
		return selectedCard
	}
}

func (c *playerController) tryTopDecking(gameState GameState, deck *deck) card.Card {
	extraCard := deck.DrawOne()
	if Playable(extraCard, gameState.CurrentColor(), gameState.LastPlayedCard()) {
		ui.Message.PlayerDrewAndPlayedCard(c.Name(), extraCard)
		return extraCard
	} else {
		c.AddCard(extraCard)
		ui.Message.PlayerPassed(c.Name())
		return nil
	}
}

func (c *playerController) selectPlayableCards(gameState GameState) []card.Card {
	currentColor := gameState.CurrentColor()
	lastPlayedCard := gameState.LastPlayedCard()

	var playableCards []card.Card
	for _, candidateCard := range c.hand {
		if Playable(candidateCard, currentColor, lastPlayedCard) {
			playableCards = append(playableCards, candidateCard)
		}
	}
	return playableCards
}

func (c *playerController) removeCardFromHand(card card.Card) error {
	cardIndex, err := c.findCardInHand(card)
	if err != nil {
		return err
	}
	c.hand[cardIndex] = c.hand[len(c.hand)-1]
	c.hand = c.hand[:len(c.hand)-1]
	return nil
}

func (c *playerController) findCardInHand(card card.Card) (int, error) {
	for index, cardInHand := range c.hand {
		if cardInHand.Equal(card) {
			return index, nil
		}
	}
	return -1, errors.New("card not found")
}
