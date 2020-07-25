package player

import (
	"github.com/antonio-muniz/uno/pkg/card"
	"github.com/antonio-muniz/uno/pkg/card/color"
	"github.com/antonio-muniz/uno/pkg/event"
	"github.com/antonio-muniz/uno/pkg/game"
	"github.com/antonio-muniz/uno/pkg/ui"
)

type humanPlayer struct {
	basicPlayer
}

func NewHumanPlayer(name string) game.Player {
	player := humanPlayer{basicPlayer: basicPlayer{name: name}}
	event.FirstCardPlayed.AddListener(player)
	event.CardPlayed.AddListener(player)
	event.ColorPicked.AddListener(player)
	return player
}

func (p humanPlayer) PickColor(gameState game.State) color.Color {
	color := ui.PromptColor()
	return color
}

func (p humanPlayer) Play(playableCards []card.Card, gameState game.State) card.Card {
	ui.Message.HumanPlayerTurnStarted(p.name)
	ui.Println(gameState)
	card := ui.PromptCardSelection(playableCards)
	return card
}

func (p humanPlayer) OnFirstCardPlayed(payload event.FirstCardPlayedPayload) {
	ui.Message.FirstCardPlayed(payload.Card)
}

func (p humanPlayer) OnCardPlayed(payload event.CardPlayedPayload) {
	ui.Message.PlayerPlayedCard(payload.PlayerName, payload.Card)
}

func (p humanPlayer) OnColorPicked(payload event.ColorPickedPayload) {
	ui.Message.PlayerPickedColor(payload.PlayerName, payload.Color)
}

func (p humanPlayer) NotifyCardsDrawn(cards []card.Card) {
	ui.Message.HumanPlayerDrewCards(cards)
}

func (p humanPlayer) NotifyNoMatchingCardsInHand(lastPlayedCard card.Card, hand []card.Card) {
	ui.Message.HumanPlayerHasNoMatchingCardsInHand(p.name, lastPlayedCard, hand)
}
