package dispatcher

import (
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

type Dispatcher struct {
	*bot.Bot
	updatesChan chan *types.Update
	nextOffset  int

	webhookServer   *http.Server
	webhookServerMu sync.RWMutex

	handlers handlers
}

type handlers struct {
	messages         []messageHandler
	callbacks        []callbackQueryHandler
	preCheckoutQuery preCheckoutQueryHandlerFunc
	shippingQuery    shippingQueryHandlerFunc
}

func NewDispatcher(token string, webhookOptions *bot.WebhookOptions) *Dispatcher {
	botInstance := bot.NewBot(token, webhookOptions)

	if webhookOptions != nil {
		res, err := botInstance.SetWebhook()
		if err != nil {
			log.Fatal("Failed to set webhook: ", err)
		}
		log.Println(res)
	}

	return &Dispatcher{
		Bot:         botInstance,
		updatesChan: make(chan *types.Update, bufferSize),
	}
}

func (d *Dispatcher) processUpdates(updatesChan <-chan *types.Update) {
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

func (d *Dispatcher) checkUpdate(update *types.Update) {
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

func (d *Dispatcher) Shutdown() {
	d.webhookServerMu.Lock()
	server := d.webhookServer
	d.webhookServerMu.Unlock()

	if server != nil {
		d.shutdownWebhookServer()
	}

	d.Cancel()

	time.Sleep(1 * time.Second)
}
