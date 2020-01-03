package card

import (
	"github.com/antonio-muniz/uno/game/card/action"
	"github.com/antonio-muniz/uno/game/card/color"
)

type WildCard struct {
	color color.Color
}

func NewWildCard() WildCard {
	return WildCard{color: color.Black}
}

func (c WildCard) Actions() []action.Action {
	return []action.Action{
		action.NewPickColorAction(),
	}
}

func (c WildCard) Color() color.Color {
	return c.color
}

func (c WildCard) String() string {
	return c.color.Paint("[(*)]")
}
