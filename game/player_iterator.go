package game

import "github.com/antonio-muniz/uno/game/ui"

const (
	left  = -1
	right = 1
)

type playerIterator struct {
	players   []*playerController
	direction int
	current   int
}

func newPlayerIterator(players []*playerController) *playerIterator {
	return &playerIterator{
		players:   players,
		direction: right,
		current:   len(players) - 1,
	}
}

func (i *playerIterator) Current() *playerController {
	return i.players[i.current]
}

func (i *playerIterator) ForEach(function func(player *playerController)) {
	for range i.players {
		function(i.Current())
		i.Next()
	}
}

func (i *playerIterator) Next() *playerController {
	playerCount := len(i.players)
	i.current = (i.current + i.direction + playerCount) % playerCount
	return i.players[i.current]
}

func (i *playerIterator) Reverse() {
	i.direction = -i.direction
	ui.Message.TurnOrderReversed()
}

func (i *playerIterator) Skip() {
	skippedPlayer := i.Next()
	ui.Message.PlayerTurnSkipped(skippedPlayer.Name())
}
