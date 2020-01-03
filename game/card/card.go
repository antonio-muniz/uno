package card

import (
	"github.com/antonio-muniz/uno/game/card/action"
	"github.com/antonio-muniz/uno/game/card/color"
)

type Card interface {
	Actions() []action.Action
	Color() color.Color
	String() string
}
