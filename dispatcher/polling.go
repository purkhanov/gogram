package dispatcher

import (
	"context"
	"errors"
	"log"
	"time"

	"github.com/purkhanov/gogram/bot"
)

const (
	pollingTimeout = 3600
	sleepTime      = 1 * time.Second
)

func (d *Dispatcher) StartPolling(skipUpdates bool) error {
	if _, err := d.Bot.DeleteWebhook(skipUpdates); err != nil {
		return err
	}

	log.Println("starting polling for updates...")

	params := bot.GetUpdatesOptions{
		Timeout: pollingTimeout,
	}

	go func() {
		defer close(d.updatesChan)

		for {
			select {
			case <-d.Ctx.Done():
				return
			default:
				params.Offset = d.nextOffset
				updates, err := d.GetUpdates(params)
				if err != nil {
					if errors.Is(err, context.Canceled) {
						return
					}

					log.Printf("error fetching updates: %v", err)
					continue
				}

				for _, update := range updates {
					d.nextOffset = update.UpdateID + 1
					select {
					case <-d.Ctx.Done():
						return
					default:
						d.updatesChan <- update
					}
				}
			}
		}
	}()

	go d.processUpdates(d.updatesChan)

	return nil
}
