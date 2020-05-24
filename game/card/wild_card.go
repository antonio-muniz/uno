package card

import (
	"github.com/antonio-muniz/uno/game/card/action"
	"github.com/antonio-muniz/uno/game/card/color"
)

type WildCard struct{}

func NewWildCard() WildCard {
	return WildCard{}
}

func (c WildCard) Actions() []action.Action {
	return []action.Action{
		action.NewPickColorAction(),
	}
}

func (c WildCard) Color() color.Color {
	return nil
}

func (c WildCard) Equal(other Card) bool {
	_, typeMatched := other.(WildCard)
	return typeMatched
}

func (c WildCard) String() string {
	return "[(*)]"
}
