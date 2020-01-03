package player

import (
	"fmt"

	"github.com/antonio-muniz/uno/game"
	"github.com/antonio-muniz/uno/game/card"
	"github.com/antonio-muniz/uno/game/card/color"
	"github.com/antonio-muniz/uno/game/ui"
)

type humanPlayer struct {
	basicPlayer
}

func NewHumanPlayer(name string) game.Player {
	return humanPlayer{basicPlayer: basicPlayer{name: name}}
}

func (p humanPlayer) PickColor(gameState game.GameState) color.Color {
	ui.Printfln(
		"Select a color: %s, %s, %s or %s?",
		color.Red,
		color.Yellow,
		color.Green,
		color.Blue,
	)

	var input string
	_, err := fmt.Scanln(&input)
	if err != nil {
		panic("PLAY THE FUCKING GAME RIGHT!")
	}

	chosenColor := color.ByName(input)
	if chosenColor == nil {
		panic("PLAY THE FUCKING GAME RIGHT!")
	}
	return chosenColor
}

func (p humanPlayer) Play(playableCards map[int]card.Card, gameState game.GameState) int {
	ui.Message.HumanPlayerTurnStarted(p.name)
	ui.Println(gameState)

	cardListing := []string{"Select a card to play:"}
	for cardID, playableCard := range playableCards {
		cardListing = append(cardListing, fmt.Sprintf("%s (enter %d)", playableCard, cardID))
	}
	ui.Printlns(cardListing)

	var input int
	_, err := fmt.Scanln(&input)
	if err != nil || playableCards[input] == nil {
		panic("PLAY THE FUCKING GAME RIGHT!")
	}

	return input
}

func (p humanPlayer) NotifyCardsDrawn(cards []card.Card) {
	ui.Message.HumanPlayerDrewCards(cards)
}

func (p humanPlayer) NotifyNoMatchingCardsInHand(lastPlayedCard card.Card, hand []card.Card) {
	ui.Message.HumanPlayerHasNoMatchingCardsInHand(p.name, lastPlayedCard, hand)
}
