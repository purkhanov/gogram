package dispatcher

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"time"

	"github.com/purkhanov/gogram/bot"
	"github.com/purkhanov/gogram/types"
)

func (d *dispatcher) StartWebhook(port uint16, params bot.SetWebhookParameters) error {
	if _, err := d.Bot.SetWebhook(params); err != nil {
		return err
	}

	if port == 0 {
		return errors.New("port can not be null")
	}

	handler := func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
			return
		}

		body, err := io.ReadAll(r.Body)
		if err != nil {
			log.Printf("unable to read request body: %v", err)
			http.Error(w, "Bad request", http.StatusBadRequest)
			return
		}
		defer r.Body.Close()

		var update types.Update

		if err := json.Unmarshal(body, &update); err != nil {
			log.Printf("unable to decode update: %v", err)
			http.Error(w, "Bad request", http.StatusBadRequest)
			return
		}

		select {
		case d.updatesChan <- &update:
			w.WriteHeader(http.StatusOK)

		case <-time.After(5 * time.Second):
			log.Printf("timeout sending update to channel")
			http.Error(w, "Service unavailable", http.StatusServiceUnavailable)

		case <-d.Ctx.Done():
			log.Printf("dispatcher stopped, cannot process update")
			http.Error(w, "Service unavailable", http.StatusServiceUnavailable)
		}
	}

	server := &http.Server{
		Addr:    fmt.Sprintf(":%d", port),
		Handler: http.HandlerFunc(handler),
	}
	d.webhookServer = server

	go func() {
		log.Printf("Starting webhook server on %d", port)
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Printf("webhook server error: %v", err)
		}
	}()

	go d.processUpdates(d.updatesChan)

	return nil
}

func (d *dispatcher) ShutdownWebhookServer() {
	if d.webhookServer != nil {
		ctx, cancel := context.WithTimeout(d.Ctx, 5*time.Second)
		defer cancel()
		d.webhookServer.Shutdown(ctx)
	}
	d.cancel()
}

func extractAddrFromURL(urlStr string) string {
	u, err := url.Parse(urlStr)
	if err != nil {
		return ""
	}
	return u.Host
}
