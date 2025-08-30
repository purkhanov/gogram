package dispatcher

import (
	"log"

	"github.com/purkhanov/gogram/commands"
	filters "github.com/purkhanov/gogram/filter"
	"github.com/purkhanov/gogram/types"
)

type messageHandlerFunc func(*types.Message) error

func (d *dispatcher) CommandHandler(msgHandler messageHandlerFunc, command commands.Command) {
	d.handlers = append(d.handlers, handler{
		filters: []filters.Filter{filters.IsCommand(command)},
		handler: func(u *types.Update) error {
			if u.Message == nil {
				return nil
			}
			return msgHandler(u.Message)
		},
	})
}

func (d *dispatcher) MessageHandler(msgHandler messageHandlerFunc, filters ...filters.Filter) {
	d.handlers = append(d.handlers, handler{
		filters: filters,
		handler: func(u *types.Update) error {
			if u.Message == nil {
				return nil
			}
			return msgHandler(u.Message)
		},
	})
}

func (d *dispatcher) processUpdates(updatesChan <-chan *types.Update) {
	for {
		select {
		case <-d.Ctx.Done():
			return

		case update, ok := <-updatesChan:
			if !ok {
				return
			}

			for _, handler := range d.handlers {
				go d.checkHandler(handler, update)
			}
		}
	}
}

func (d *dispatcher) checkHandler(handler handler, update *types.Update) {
	for _, filter := range handler.filters {
		if !filter(update.Message) {
			return
		}
	}

	if err := handler.handler(update); err != nil {
		log.Printf("error handling update: %v", err)
	}
}
