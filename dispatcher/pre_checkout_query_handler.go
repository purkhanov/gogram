package dispatcher

import (
	"github.com/purkhanov/gogram/types"
)

type preCheckoutQueryHandlerFunc func(*types.PreCheckoutQuery)

func (d *dispatcher) OnPreCheckoutQuery(handler preCheckoutQueryHandlerFunc) {
	d.handlers.preCheckoutQuery = handler
}

func (d *dispatcher) handlePreCheckoutQuery(preCheckoutQuery *types.PreCheckoutQuery) {
	d.handlers.preCheckoutQuery(preCheckoutQuery)
}
