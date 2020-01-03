package game

import (
	"github.com/antonio-muniz/uno/game/card"
	"github.com/antonio-muniz/uno/game/card/color"
)

type Player interface {
	Name() string
	PickColor(GameState) color.Color
	Play(map[int]card.Card, GameState) int
	NotifyCardsDrawn([]card.Card)
	NotifyNoMatchingCardsInHand(card.Card, []card.Card)
}
