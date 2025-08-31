package dispatcher

import (
	"log"

	"github.com/purkhanov/gogram/types"
)

type preCheckoutQueryHandlerFunc func(*types.PreCheckoutQuery) error

func (d *dispatcher) PreCheckoutQueryHandler(preCheckoutHandler preCheckoutQueryHandlerFunc) {
	d.handlers.preCheckoutQueryHandler = preCheckoutHandler
}

func (d *dispatcher) preCheckoutQueryHandler(preCheckoutQuery *types.PreCheckoutQuery) {
	err := d.handlers.preCheckoutQueryHandler(preCheckoutQuery)
	if err != nil {
		log.Printf("error handling preCheckoutQuery: %v", err)
	}
}
