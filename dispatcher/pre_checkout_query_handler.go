package dispatcher

import (
	"github.com/purkhanov/gogram/types"
)

type preCheckoutQueryHandlerFunc func(*types.PreCheckoutQuery)

func (d *Dispatcher) OnPreCheckoutQuery(handler preCheckoutQueryHandlerFunc) {
	d.handlers.preCheckoutQuery = handler
}

func (d *Dispatcher) handlePreCheckoutQuery(preCheckoutQuery *types.PreCheckoutQuery) {
	d.handlers.preCheckoutQuery(preCheckoutQuery)
}
