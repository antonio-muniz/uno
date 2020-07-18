package game

import (
	"github.com/antonio-muniz/uno/pkg/card"
	"github.com/antonio-muniz/uno/pkg/card/color"
)

type Player interface {
	Name() string
	PickColor(gameState State) color.Color
	Play(playableCards []card.Card, gameState State) card.Card
	NotifyCardsDrawn(drawnCards []card.Card)
	NotifyNoMatchingCardsInHand(lastPlayedCard card.Card, hand []card.Card)
}
