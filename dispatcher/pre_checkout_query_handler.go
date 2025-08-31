package dispatcher

import (
	"github.com/purkhanov/gogram/types"
)

type preCheckoutQueryHandlerFunc func(*types.PreCheckoutQuery)

func (d *dispatcher) PreCheckoutQueryHandler(preCheckoutHandler preCheckoutQueryHandlerFunc) {
	d.handlers.preCheckoutQueryHandler = preCheckoutHandler
}

func (d *dispatcher) callPreCheckoutQueryHandler(preCheckoutQuery *types.PreCheckoutQuery) {
	d.handlers.preCheckoutQueryHandler(preCheckoutQuery)
}
