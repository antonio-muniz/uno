package card

import (
	"github.com/antonio-muniz/uno/game/card/action"
	"github.com/antonio-muniz/uno/game/card/color"
)

type WildDrawFourCard struct{}

func NewWildDrawFourCard() WildDrawFourCard {
	return WildDrawFourCard{}
}

func (c WildDrawFourCard) Actions() []action.Action {
	return []action.Action{
		action.NewPickColorAction(),
		action.NewSkipTurnAction(),
		action.NewDrawCardsAction(4),
	}
}

func (c WildDrawFourCard) Color() color.Color {
	return nil
}

func (c WildDrawFourCard) Equal(other Card) bool {
	_, typeMatched := other.(WildDrawFourCard)
	return typeMatched
}

func (c WildDrawFourCard) String() string {
	return "[+4!]"
}
