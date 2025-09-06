package dispatcher

import "github.com/purkhanov/gogram/types"

type shippingQueryHandlerFunc func(*types.ShippingQuery)

func (d *Dispatcher) OnShippingQuery(handler shippingQueryHandlerFunc) {
	d.handlers.shippingQuery = handler
}

func (d *Dispatcher) handleShippingQuery(shippingQuery *types.ShippingQuery) {
	d.handlers.shippingQuery(shippingQuery)
}
