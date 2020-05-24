package game

import (
	"github.com/antonio-muniz/uno/game/card"
	"github.com/antonio-muniz/uno/game/card/color"
)

type Player interface {
	Name() string
	PickColor(gameState GameState) color.Color
	Play(playableCards []card.Card, gameState GameState) card.Card
	NotifyCardsDrawn(drawnCards []card.Card)
	NotifyNoMatchingCardsInHand(lastPlayedCard card.Card, hand []card.Card)
}
