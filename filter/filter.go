package filters

import "github.com/purkhanov/gogram/types"

type Filter func(*types.Message) bool

func TextEquals(text string) Filter {
	return func(m *types.Message) bool {
		return m.Text == text
	}
}
