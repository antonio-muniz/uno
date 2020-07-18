package game

import (
	"math/rand"

	"github.com/antonio-muniz/uno/pkg/card"
	"github.com/antonio-muniz/uno/pkg/card/color"
)

func createUnoDeck() *deck {
	deck := &deck{}
	fillDeck(deck)
	return deck
}

type deck struct {
	cards []card.Card
}

func (d *deck) DrawOne() card.Card {
	return d.Draw(1)[0]
}

func (d *deck) Draw(amount int) []card.Card {
	if len(d.cards) < amount {
		fillDeck(d)
	}
	cards := d.cards[0:amount]
	d.cards = d.cards[amount:]
	return cards
}

func fillDeck(deck *deck) {
	cards := make([]card.Card, 0, 108)

	cards = append(cards, createBlackCards()...)
	cards = append(cards, createColorCards(color.Red)...)
	cards = append(cards, createColorCards(color.Yellow)...)
	cards = append(cards, createColorCards(color.Green)...)
	cards = append(cards, createColorCards(color.Blue)...)

	shuffleCards(cards)

	deck.cards = cards
}

func createColorCards(cardColor color.Color) []card.Card {
	zeroCard := card.NewNumberCard(cardColor, 0)
	skipCard := card.NewSkipCard(cardColor)
	reverseCard := card.NewReverseCard(cardColor)
	drawTwoCard := card.NewDrawTwoCard(cardColor)

	cards := []card.Card{
		zeroCard,
		skipCard, skipCard,
		reverseCard, reverseCard,
		drawTwoCard, drawTwoCard,
	}

	for n := 1; n <= 9; n++ {
		numberCard := card.NewNumberCard(cardColor, n)
		cards = append(cards, numberCard, numberCard)
	}

	return cards
}

func createBlackCards() []card.Card {
	wildCard := card.NewWildCard()
	wildDrawFourCard := card.NewWildDrawFourCard()

	return []card.Card{
		wildCard, wildCard, wildCard, wildCard,
		wildDrawFourCard, wildDrawFourCard, wildDrawFourCard, wildDrawFourCard,
	}
}

func shuffleCards(cards []card.Card) {
	rand.Shuffle(len(cards), func(i, j int) { cards[i], cards[j] = cards[j], cards[i] })
}
