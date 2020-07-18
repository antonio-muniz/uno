package player

import (
	"math/rand"

	"github.com/antonio-muniz/uno/pkg/card"
	"github.com/antonio-muniz/uno/pkg/card/color"
	"github.com/antonio-muniz/uno/pkg/game"
)

type naivePlayer struct {
	basicPlayer
}

func NewNaivePlayer(name string) game.Player {
	return naivePlayer{basicPlayer: basicPlayer{name: name}}
}

func (p naivePlayer) PickColor(gameState game.State) color.Color {
	return []color.Color{
		color.Red,
		color.Yellow,
		color.Blue,
		color.Green,
	}[rand.Intn(4)]
}

func (p naivePlayer) Play(playableCards []card.Card, gameState game.State) card.Card {
	for _, card := range playableCards {
		return card
	}
	return nil
}
