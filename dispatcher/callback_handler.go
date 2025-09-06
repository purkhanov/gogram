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

func (d *Dispatcher) OnCallbackQuery(handler callbackQueryHandlerFunc, filters ...filters.CallbackFilter) {
	d.handlers.callbacks = append(d.handlers.callbacks, callbackQueryHandler{
		filters: filters,
		handler: handler,
	})
}

func (d *Dispatcher) handleCallbackQuery(callbackQuery *types.CallbackQuery) {
	for _, handler := range d.handlers.callbacks {
		if !matchesFilters(handler.filters, callbackQuery) {
			continue
		}

		handler.handler(callbackQuery)
	}
}

func matchesFilters(filters []filters.CallbackFilter, callbackQuery *types.CallbackQuery) bool {
	for _, filter := range filters {
		if !filter(callbackQuery) {
			return false
		}
	}

	return true
}
