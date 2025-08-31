package dispatcher

import (
	filters "github.com/purkhanov/gogram/filter"
	"github.com/purkhanov/gogram/types"
)

type callbackQueryHandlerFunc func(*types.CallbackQuery)

type callbackQueryHandler struct {
	filters []filters.CallbackFilter
	handler callbackQueryHandlerFunc
}

func (d *dispatcher) CallbackQueryHandler(cbHandler callbackQueryHandlerFunc, filters ...filters.CallbackFilter) {
	d.handlers.callbackHandler = append(d.handlers.callbackHandler, callbackQueryHandler{
		filters: filters,
		handler: cbHandler,
	})
}

func (d *dispatcher) callbackQueryHandler(callback *types.CallbackQuery) {
	for _, cbHandler := range d.handlers.callbackHandler {
		matches := true

		for _, filter := range cbHandler.filters {
			if !filter(callback) {
				matches = false
				break
			}
		}

		if !matches {
			continue
		}

		cbHandler.handler(callback)
	}
}
