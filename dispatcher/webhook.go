package dispatcher

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
	"time"

	"github.com/purkhanov/gogram/bot"
	"github.com/purkhanov/gogram/types"
)

func (d *Dispatcher) SetupWebhook(webhookOptions bot.WebhookOptions) error {
	_, err := d.Bot.SetWebhook(webhookOptions)

	return err
}

func (d *Dispatcher) StartWebhookServer(port uint16) error {
	if port == 0 {
		return errors.New("port cannot be zero")
	}

	handler := func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
			return
		}

		r.Body = http.MaxBytesReader(w, r.Body, maxRequestBodySize)

		body, err := io.ReadAll(r.Body)
		if err != nil {
			log.Printf("Unable to read request body: %v", err)
			http.Error(w, "Bad request", http.StatusBadRequest)
			return
		}
		defer r.Body.Close()

		var update types.Update

		if err := json.Unmarshal(body, &update); err != nil {
			log.Printf("Unable to decode update: %v", err)
			http.Error(w, "Bad request", http.StatusBadRequest)
			return
		}

		select {
		case d.updatesChan <- &update:
			w.WriteHeader(http.StatusOK)
			w.Write([]byte("OK"))

		case <-time.After(channelTimeout):
			log.Printf("Timeout sending update to channel")
			http.Error(w, "Service unavailable", http.StatusServiceUnavailable)

		case <-d.ctx.Done():
			log.Printf("Dispatcher stopped, cannot process update")
			http.Error(w, "Service unavailable", http.StatusServiceUnavailable)
		}
	}

	server := &http.Server{
		Addr:         fmt.Sprintf(":%d", port),
		Handler:      http.HandlerFunc(handler),
		ReadTimeout:  readTimeout,
		WriteTimeout: writeTimeout,
	}

	d.webhookServerMu.Lock()
	d.webhookServer = server
	d.webhookServerMu.Unlock()

	go func() {
		log.Printf("Starting webhook server on port %d", port)
		if err := server.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			log.Printf("Webhook server error: %v", err)
		}
	}()

	go d.processUpdates(d.updatesChan)

	return nil
}

func (d *Dispatcher) shutdownWebhookServer() {
	d.webhookServerMu.Lock()
	server := d.webhookServer
	d.webhookServer = nil
	d.webhookServerMu.Unlock()

	if server != nil {
		log.Println("ShutdownWebhookServer: shutting down HTTP server")

		ctx, cancel := context.WithTimeout(d.ctx, shutdownTimeout)
		defer cancel()

		if err := server.Shutdown(ctx); err != nil {
			log.Printf("ShutdownWebhookServer: error: %v", err)
		} else {
			log.Println("ShutdownWebhookServer: HTTP server stopped")
		}
	}
}

// with gin freamwork
type ginContext interface {
	ShouldBindJSON(obj any) error
	JSON(code int, obj any)
}

func (d *Dispatcher) GinWebhookHandler(ctx ginContext) {
	var update types.Update

	if err := ctx.ShouldBindJSON(&update); err != nil {
		ctx.JSON(http.StatusBadRequest, map[string]any{
			"error":   "Invalid JSON",
			"details": err.Error(),
		})

		return
	}

	go d.checkUpdate(&update)

	ctx.JSON(http.StatusOK, map[string]string{"status": "accepted"})

}
