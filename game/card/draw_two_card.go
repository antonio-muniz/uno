package card

import (
	"github.com/antonio-muniz/uno/game/card/action"
	"github.com/antonio-muniz/uno/game/card/color"
)

type DrawTwoCard struct {
	color color.Color
}

func NewDrawTwoCard(color color.Color) DrawTwoCard {
	return DrawTwoCard{color: color}
}

func (c DrawTwoCard) Actions() []action.Action {
	return []action.Action{
		action.NewSkipTurnAction(),
		action.NewDrawCardsAction(2),
	}
}

func (c DrawTwoCard) Color() color.Color {
	return c.color
}

func (c DrawTwoCard) String() string {
	return c.color.Paint("[+2!]")
}
