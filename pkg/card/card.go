package card

import (
	"github.com/antonio-muniz/uno/pkg/card/action"
	"github.com/antonio-muniz/uno/pkg/card/color"
)

type Card interface {
	Actions() []action.Action
	Color() color.Color
	Equal(other Card) bool
	String() string
}
