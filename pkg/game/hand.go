package game

import (
	"github.com/antonio-muniz/uno/pkg/card"
)

type Hand struct {
	cards []card.Card
}

func NewHand() *Hand {
	return &Hand{cards: make([]card.Card, 0, 7)}
}

func (h *Hand) AddCards(cards []card.Card) {
	h.cards = append(h.cards, cards...)
}

func (h *Hand) Cards() []card.Card {
	cards := make([]card.Card, len(h.cards))
	copy(cards, h.cards)
	return cards
}

func (h *Hand) Empty() bool {
	return len(h.cards) == 0
}

func (h *Hand) PlayableCards(lastPlayedCard card.Card) []card.Card {
	var playableCards []card.Card
	for _, candidateCard := range h.cards {
		if Playable(candidateCard, lastPlayedCard) {
			playableCards = append(playableCards, candidateCard)
		}
	}
	return playableCards
}

func (h *Hand) RemoveCard(card card.Card) {
	for index, cardInHand := range h.cards {
		if cardInHand.Equal(card) {
			h.cards[index] = h.cards[len(h.cards)-1]
			h.cards = h.cards[:len(h.cards)-1]
			return
		}
	}
}

func (h *Hand) Size() int {
	return len(h.cards)
}
