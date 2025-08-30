package dispatcher

import (
	"context"
	"net/http"

	"github.com/purkhanov/gogram/bot"
	filters "github.com/purkhanov/gogram/filter"
	"github.com/purkhanov/gogram/types"
)

const bufferSize = 20

type dispatcher struct {
	Bot         *bot.Bot
	updatesChan chan *types.Update
	nextOffset  int

	webhookServer *http.Server

	Ctx      context.Context
	cancel   context.CancelFunc
	handlers []handler
}

type handler struct {
	filters []filters.Filter
	handler func(*types.Update) error
}

func NewDispatcher(bot *bot.Bot) *dispatcher {
	ctx, cancel := context.WithCancel(context.Background())

	return &dispatcher{
		Bot:         bot,
		updatesChan: make(chan *types.Update, bufferSize),
		Ctx:         ctx,
		cancel:      cancel,
	}
}
