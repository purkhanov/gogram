package dispatcher

import (
	"context"
	"log"
	"net/http"
	"sync"
	"time"

	"github.com/purkhanov/gogram/bot"
	"github.com/purkhanov/gogram/types"
)

const (
	bufferSize         = 20
	readTimeout        = 10 * time.Second
	writeTimeout       = 10 * time.Second
	shutdownTimeout    = 5 * time.Second
	channelTimeout     = 5 * time.Second
	maxRequestBodySize = 1 << 20 // 1 MB
)

type dispatcher struct {
	Bot            *bot.Bot
	WebhookOptions bot.WebhookOptions
	updatesChan    chan *types.Update
	nextOffset     int

	webhookServer   *http.Server
	webhookServerMu sync.RWMutex

	Ctx    context.Context
	cancel context.CancelFunc

	handlers handlers
}

type handlers struct {
	messages         []messageHandler
	callbacks        []callbackQueryHandler
	preCheckoutQuery preCheckoutQueryHandlerFunc
	shippingQuery    shippingQueryHandlerFunc
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
		d.handleMessage(update.Message)

	case update.CallbackQuery != nil:
		d.handleCallbackQuery(update.CallbackQuery)

	case update.PreCheckoutQuery != nil:
		d.handlePreCheckoutQuery(update.PreCheckoutQuery)

	case update.ShippingQuery != nil:
		d.handleShippingQuery(update.ShippingQuery)

	default:
		log.Println("unknown update type", update)
	}
}
