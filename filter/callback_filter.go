package filters

import (
	"github.com/purkhanov/gogram/types"
)

type CallbackFilter func(*types.CallbackQuery) bool

func CallbackDataEquals(data string) CallbackFilter {
	return func(cb *types.CallbackQuery) bool {
		return data == cb.Data
	}
}
