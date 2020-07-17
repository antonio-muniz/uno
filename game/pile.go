package game

import (
	"github.com/antonio-muniz/uno/game/card"
	"github.com/antonio-muniz/uno/game/card/color"
)

type pile struct {
	cards        []card.Card
	currentColor color.Color
}

func createEmptyPile() *pile {
	return &pile{cards: make([]card.Card, 0, 54)}
}

func (p *pile) Add(card card.Card) {
	p.cards = append(p.cards, card)
	p.currentColor = card.Color()
}

func (p *pile) CurrentColor() color.Color {
	return p.currentColor
}

func (p *pile) ReplaceTop(card card.Card) {
	p.cards[len(p.cards)-1] = card
	p.currentColor = card.Color()
}

func (p *pile) SetCurrentColor(color color.Color) {
	p.currentColor = color
}

func (p *pile) Top() card.Card {
	pileSize := len(p.cards)
	if pileSize == 0 {
		return nil
	}
	return p.cards[pileSize-1]
}
