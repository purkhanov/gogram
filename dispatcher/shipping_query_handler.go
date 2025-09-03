package dispatcher

import "github.com/purkhanov/gogram/types"

type shippingQueryHandlerFunc func(*types.ShippingQuery)

func (d *dispatcher) OnShippingQuery(handler shippingQueryHandlerFunc) {
	d.handlers.shippingQuery = handler
}

func (d *dispatcher) handleShippingQuery(shippingQuery *types.ShippingQuery) {
	d.handlers.shippingQuery(shippingQuery)
}
