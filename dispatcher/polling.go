package dispatcher

import (
	"log"
	"time"

	"github.com/purkhanov/gogram/bot"
)

const (
	timeout   = 3600
	sleepTime = 1 * time.Second
)

func (d *dispatcher) StartPolling(skipUpdates bool) error {
	if _, err := d.Bot.DeleteWebhook(skipUpdates); err != nil {
		return err
	}

	log.Println("starting polling for updates...")

	params := bot.GetUpdatesOptions{
		Timeout: timeout,
	}

	go func() {
		defer close(d.updatesChan)

		for {
			select {
			case <-d.Ctx.Done():
				log.Println("stopping polling")
				return
			default:
				params.Offset = d.nextOffset
				updates, err := d.Bot.GetUpdates(params)
				if err != nil {
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

func (d *dispatcher) Shutdown() {
	d.cancel()
}
