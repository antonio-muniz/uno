package game

import (
	"github.com/antonio-muniz/uno/pkg/card"
)

type Pile struct {
	cards []card.Card
}

func NewPile() *Pile {
	return &Pile{cards: make([]card.Card, 0, 54)}
}

func (p *Pile) Add(card card.Card) {
	p.cards = append(p.cards, card)
}

func (p *Pile) Cards() []card.Card {
	cards := make([]card.Card, len(p.cards))
	copy(cards, p.cards)
	return cards
}

func (p *Pile) ReplaceTop(card card.Card) {
	p.cards[len(p.cards)-1] = card
}

func (p *Pile) Top() card.Card {
	pileSize := len(p.cards)
	if pileSize == 0 {
		return nil
	}
	return p.cards[pileSize-1]
}
