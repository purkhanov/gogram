package dispatcher

import (
	"github.com/purkhanov/gogram/commands"
	filters "github.com/purkhanov/gogram/filter"
	"github.com/purkhanov/gogram/types"
)

type messageHandlerFunc func(*types.Message)

type messageHandler struct {
	filters []filters.MessageFilter
	handler messageHandlerFunc
}

func (d *Dispatcher) OnCommand(command commands.Command, handler messageHandlerFunc) {
	d.OnMessage(handler, filters.IsCommand(command))
}

func (d *Dispatcher) OnMessage(handler messageHandlerFunc, filters ...filters.MessageFilter) {
	d.handlers.messages = append(d.handlers.messages, messageHandler{
		filters: filters,
		handler: handler,
	})
}

func (d *Dispatcher) handleMessage(msg *types.Message) {
	for _, message := range d.handlers.messages {
		matches := true

		for _, filter := range message.filters {
			if !filter(msg) {
				matches = false
				break
			}
		}

		if !matches {
			continue
		}

		message.handler(msg)
	}
}
