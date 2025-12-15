package dispatcher

import (
	"context"
	"errors"
	"log"

	"github.com/purkhanov/gogram/bot"
)

func (d *Dispatcher) StartPolling(skipUpdates bool) error {
	if _, err := d.Bot.DeleteWebhook(skipUpdates); err != nil {
		return err
	}

	log.Println("starting polling for updates...")

	go func() {
		defer close(d.updatesChan)

		for {
			select {
			case <-d.ctx.Done():
				return
			default:
				params := bot.GetUpdatesOptions{Offset: d.nextOffset}
				updates, err := d.Bot.GetUpdates(params)
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
					case <-d.ctx.Done():
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
