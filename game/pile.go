package game

import (
	"github.com/antonio-muniz/uno/game/card"
)

type pile struct {
	cards []card.Card
}

func createEmptyPile() *pile {
	return &pile{cards: make([]card.Card, 0, 54)}
}

func (p *pile) Add(card card.Card) {
	p.cards = append(p.cards, card)
}

func (p *pile) ReplaceTop(card card.Card) {
	p.cards[len(p.cards)-1] = card
}

func (p *pile) Top() card.Card {
	pileSize := len(p.cards)
	if pileSize == 0 {
		return nil
	}
	return p.cards[pileSize-1]
}
