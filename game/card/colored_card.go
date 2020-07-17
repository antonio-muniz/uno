package card

import (
	"github.com/antonio-muniz/uno/game/card/action"
	"github.com/antonio-muniz/uno/game/card/color"
)

type ColoredCard struct {
	card  Card
	color color.Color
}

func NewColoredCard(card Card, color color.Color) ColoredCard {
	return ColoredCard{
		card:  card,
		color: color,
	}
}

func (c ColoredCard) Actions() []action.Action {
	return c.Actions()
}

func (c ColoredCard) Color() color.Color {
	return c.color
}

func (c ColoredCard) Equal(other Card) bool {
	return c.card.Equal(other) && c.color == other.Color()
}

func (c ColoredCard) SameType(other Card) bool {
	return c.card.SameType(other)
}

func (c ColoredCard) String() string {
	return c.color.Paint(c.card.String())
}
