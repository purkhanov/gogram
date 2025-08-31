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

func (d *dispatcher) CommandHandler(msgHandler messageHandlerFunc, command commands.Command) {
	d.handlers.messageHandler = append(d.handlers.messageHandler, messageHandler{
		filters: []filters.MessageFilter{filters.IsCommand(command)},
		handler: msgHandler,
	})
}

func (d *dispatcher) MessageHandler(msgHandler messageHandlerFunc, filters ...filters.MessageFilter) {
	d.handlers.messageHandler = append(d.handlers.messageHandler, messageHandler{
		filters: filters,
		handler: msgHandler,
	})
}

func (d *dispatcher) messageHandler(msg *types.Message) {
	for _, msgHandler := range d.handlers.messageHandler {
		matches := true

		for _, filter := range msgHandler.filters {
			if !filter(msg) {
				matches = false
				break
			}
		}

		if !matches {
			continue
		}

		msgHandler.handler(msg)
	}
}
