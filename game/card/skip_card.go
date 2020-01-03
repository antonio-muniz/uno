package card

import (
	"github.com/antonio-muniz/uno/game/card/action"
	"github.com/antonio-muniz/uno/game/card/color"
)

type SkipCard struct {
	color color.Color
}

func NewSkipCard(color color.Color) SkipCard {
	return SkipCard{color: color}
}

func (c SkipCard) Actions() []action.Action {
	return []action.Action{
		action.NewSkipTurnAction(),
	}
}

func (c SkipCard) Color() color.Color {
	return c.color
}

func (c SkipCard) String() string {
	return c.color.Paint("[(/)]")
}
