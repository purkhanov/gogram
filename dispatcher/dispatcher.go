package dispatcher

import (
	"context"
	"net/http"

	"github.com/purkhanov/gogram/bot"
	"github.com/purkhanov/gogram/types"
)

const bufferSize = 20

type dispatcher struct {
	Bot         *bot.Bot
	updatesChan chan *types.Update
	nextOffset  int

	webhookServer *http.Server

	Ctx    context.Context
	cancel context.CancelFunc

	handlers handlers
}

type handlers struct {
	messageHandler          []messageHandler
	callbackHandler         []callbackQueryHandler
	preCheckoutQueryHandler preCheckoutQueryHandlerFunc
}

func NewDispatcher(bot *bot.Bot) *dispatcher {
	ctx, cancel := context.WithCancel(bot.Ctx)

	return &dispatcher{
		Bot:         bot,
		updatesChan: make(chan *types.Update, bufferSize),
		Ctx:         ctx,
		cancel:      cancel,
	}
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

			go d.checkUpdate(update)
		}
	}
}

func (d *dispatcher) checkUpdate(update *types.Update) {
	switch {
	case update.Message != nil:
		d.messageHandler(update.Message)
	case update.CallbackQuery != nil:
		d.callbackQueryHandler(update.CallbackQuery)
	case update.PreCheckoutQuery != nil:
		d.callPreCheckoutQueryHandler(update.PreCheckoutQuery)
	}
}
