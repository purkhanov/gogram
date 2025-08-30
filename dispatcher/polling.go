package dispatcher

import (
	"log"
	"time"

	"github.com/purkhanov/gogram/bot"
	"github.com/purkhanov/gogram/types"
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

	params := bot.GetUpdateParams{
		Timeout: timeout,
	}

	updatesChan := make(chan *types.Update, d.bufferSize)

	go func() {
		defer close(updatesChan)

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

					select {
					case <-d.Ctx.Done():
						return
					case <-time.After(sleepTime):
					}
					continue
				}

				if len(updates) == 0 {
					select {
					case <-d.Ctx.Done():
						return
					case <-time.After(100 * time.Millisecond):
					}
					continue
				}

				for _, update := range updates {
					d.nextOffset = update.UpdateID + 1
					select {
					case updatesChan <- update:
					case <-d.Ctx.Done():
						return
					}
				}
			}
		}
	}()

	go d.processUpdates(updatesChan)

	return nil
}

func (d *dispatcher) Shutdown() {
	d.cancel()
}
