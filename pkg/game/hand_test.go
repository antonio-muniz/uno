package game_test

import (
	"testing"

	"github.com/antonio-muniz/uno/pkg/card"
	"github.com/antonio-muniz/uno/pkg/card/color"
	"github.com/antonio-muniz/uno/pkg/game"
	"github.com/stretchr/testify/require"
)

func TestAddCards(t *testing.T) {
	hand := game.NewHand()
	hand.AddCards([]card.Card{
		card.NewNumberCard(color.Blue, 7),
		card.NewWildCard(),
	})
	require.ElementsMatch(t, []card.Card{
		card.NewNumberCard(color.Blue, 7),
		card.NewWildCard(),
	}, hand.Cards())
}

func TestEmpty(t *testing.T) {
	hand := game.NewHand()
	require.True(t, hand.Empty())
	hand.AddCards([]card.Card{
		card.NewNumberCard(color.Blue, 7),
		card.NewWildCard(),
	})
	require.False(t, hand.Empty())
}

func TestPlayableCards(t *testing.T) {
	hand := game.NewHand()
	hand.AddCards([]card.Card{
		card.NewNumberCard(color.Blue, 5),
		card.NewNumberCard(color.Green, 8),
		card.NewNumberCard(color.Green, 7),
		card.NewWildCard(),
		card.NewReverseCard(color.Yellow),
		card.NewDrawTwoCard(color.Blue),
	})
	lastPlayedCard := card.NewNumberCard(color.Blue, 7)
	playableCards := hand.PlayableCards(lastPlayedCard)
	require.ElementsMatch(t, []card.Card{
		card.NewNumberCard(color.Blue, 5),
		card.NewNumberCard(color.Green, 7),
		card.NewWildCard(),
		card.NewDrawTwoCard(color.Blue),
	}, playableCards)
}

func TestRemoveCard(t *testing.T) {
	hand := game.NewHand()
	hand.AddCards([]card.Card{
		card.NewWildCard(),
		card.NewReverseCard(color.Yellow),
		card.NewDrawTwoCard(color.Blue),
	})

	result := hand.RemoveCard(card.NewReverseCard(color.Yellow))
	require.True(t, result)
	require.Equal(t, []card.Card{
		card.NewWildCard(),
		card.NewDrawTwoCard(color.Blue),
	}, hand.Cards())

	result = hand.RemoveCard(card.NewDrawTwoCard(color.Red))
	require.False(t, result)
	require.Equal(t, []card.Card{
		card.NewWildCard(),
		card.NewDrawTwoCard(color.Blue),
	}, hand.Cards())
}

func TestSize(t *testing.T) {
	hand := game.NewHand()
	require.Equal(t, 0, hand.Size())
	hand.AddCards([]card.Card{
		card.NewNumberCard(color.Green, 7),
		card.NewWildCard(),
		card.NewReverseCard(color.Yellow),
	})
	require.Equal(t, 3, hand.Size())
}
