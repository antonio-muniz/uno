package game

import (
	"github.com/antonio-muniz/uno/game/card"
	"github.com/antonio-muniz/uno/game/card/color"
)

func Playable(candidateCard card.Card, currentColor color.Color, lastPlayedCard card.Card) bool {
	if candidateCard.Color() == currentColor {
		return true
	}

	switch candidateCard := candidateCard.(type) {
	case card.WildCard, card.WildDrawFourCard:
		return true
	case card.DrawTwoCard:
		_, isDrawTwoCard := lastPlayedCard.(card.DrawTwoCard)
		return isDrawTwoCard
	case card.ReverseCard:
		_, isReverseCard := lastPlayedCard.(card.ReverseCard)
		return isReverseCard
	case card.SkipCard:
		_, isSkipCard := lastPlayedCard.(card.SkipCard)
		return isSkipCard
	case card.NumberCard:
		lastPlayedCard, isNumberCard := lastPlayedCard.(card.NumberCard)
		return isNumberCard && lastPlayedCard.Number() == candidateCard.Number()
	default:
		return false
	}
}
