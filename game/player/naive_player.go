package player

import (
	"math/rand"
	"time"

	"github.com/antonio-muniz/uno/game/card"
	"github.com/antonio-muniz/uno/game/card/color"
	"github.com/antonio-muniz/uno/game"
)

type naivePlayer struct {
	basicPlayer
}

func NewNaivePlayer(name string) game.Player {
	return naivePlayer{basicPlayer: basicPlayer{name: name}}
}

func (p naivePlayer) PickColor(gameState game.GameState) color.Color {
	rand.Seed(time.Now().UnixNano())
	return []color.Color{
		color.Red,
		color.Yellow,
		color.Blue,
		color.Green,
	}[rand.Intn(4)]
}

func (p naivePlayer) Play(playableCards map[int]card.Card, gameState game.GameState) int {
	for cardID := range playableCards {
		return cardID
	}
	return -1
}
