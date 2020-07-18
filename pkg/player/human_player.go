package player

import (
	"github.com/antonio-muniz/uno/pkg"
	"github.com/antonio-muniz/uno/pkg/card"
	"github.com/antonio-muniz/uno/pkg/card/color"
	"github.com/antonio-muniz/uno/pkg/ui"
)

type humanPlayer struct {
	basicPlayer
}

func NewHumanPlayer(name string) game.Player {
	return humanPlayer{basicPlayer: basicPlayer{name: name}}
}

func (p humanPlayer) PickColor(gameState game.GameState) color.Color {
	return ui.PromptColor()
}

func (p humanPlayer) Play(playableCards []card.Card, gameState game.GameState) card.Card {
	ui.Message.HumanPlayerTurnStarted(p.name)
	ui.Println(gameState)

	return ui.PromptCardSelection(playableCards)
}

func (p humanPlayer) NotifyCardsDrawn(cards []card.Card) {
	ui.Message.HumanPlayerDrewCards(cards)
}

func (p humanPlayer) NotifyNoMatchingCardsInHand(lastPlayedCard card.Card, hand []card.Card) {
	ui.Message.HumanPlayerHasNoMatchingCardsInHand(p.name, lastPlayedCard, hand)
}