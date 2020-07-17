package player

import (
	"github.com/antonio-muniz/uno/game"
	"github.com/antonio-muniz/uno/game/card"
	"github.com/antonio-muniz/uno/game/card/color"
)

type goodPlayer struct {
	basicPlayer
}

func NewGoodPlayer(name string) game.Player {
	return goodPlayer{basicPlayer: basicPlayer{name: name}}
}

func (p goodPlayer) PickColor(gameState game.GameState) color.Color {
	if len(gameState.CurrentPlayerHand()) == 0 {
		return color.Blue
	}

	colorCounts := make(map[color.Color]int)
	for _, card := range gameState.CurrentPlayerHand() {
		if card.Color() == nil {
			colorCounts[color.Blue]++
			colorCounts[color.Green]++
			colorCounts[color.Red]++
			colorCounts[color.Yellow]++
		} else {
			colorCounts[card.Color()]++
		}
	}

	var (
		mostFrequentColor       color.Color
		mostFrequentColorAmount int
	)
	for availableColor, amount := range colorCounts {
		if amount > mostFrequentColorAmount {
			mostFrequentColorAmount = amount
			mostFrequentColor = availableColor
		}
	}

	return mostFrequentColor
}

func (p goodPlayer) Play(playableCards []card.Card, gameState game.GameState) card.Card {
	mostDiscardableCardIndex := 0
	maxSpareCards := 0

	for cardIndex, playableCard := range playableCards {
		spareCards := 0
		for _, handCard := range gameState.CurrentPlayerHand() {
			if game.Playable(handCard, playableCard) {
				spareCards++
			}
		}
		if spareCards > maxSpareCards {
			maxSpareCards = spareCards
			mostDiscardableCardIndex = cardIndex
		}
	}

	return playableCards[mostDiscardableCardIndex]
}
