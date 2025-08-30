package dispatcher

import (
	"context"

	"github.com/purkhanov/gogram/bot"
	filters "github.com/purkhanov/gogram/filter"
	"github.com/purkhanov/gogram/types"
)

type dispatcher struct {
	Bot        *bot.Bot
	bufferSize uint8
	nextOffset int

	Ctx      context.Context
	cancel   context.CancelFunc
	handlers []handler
	// errorFunc func(error, *types.Update)
}

type handler struct {
	filters []filters.Filter
	handler func(*types.Update) error
}

func NewDispatcher(bot *bot.Bot) *dispatcher {
	ctx, cancel := context.WithCancel(context.Background())

	return &dispatcher{
		Bot:        bot,
		bufferSize: 20,
		Ctx:        ctx,
		cancel:     cancel,
	}
}
