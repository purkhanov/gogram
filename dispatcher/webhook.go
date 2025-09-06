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

	"github.com/purkhanov/gogram/types"
)

func (d *dispatcher) SetupWebhook() error {
	_, err := d.Bot.SetWebhook(d.WebhookOptions)

	return err
}

func (d *dispatcher) StartWebhookServer(port uint16) error {
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

		case <-d.Ctx.Done():
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

func (d *dispatcher) ShutdownWebhookServer() {
	d.webhookServerMu.Lock()
	server := d.webhookServer
	d.webhookServerMu.Unlock()

	if server != nil {
		ctx, cancel := context.WithTimeout(d.Ctx, shutdownTimeout)
		defer cancel()

		if err := server.Shutdown(ctx); err != nil {
			log.Printf("Error shutting down webhook server: %v", err)
		}
	}

	d.cancel()
}

// with gin freamwork
type ginContext interface {
	ShouldBindJSON(obj any) error
	JSON(code int, obj any)
}

func (d *dispatcher) GinWebhookHandler(ctx ginContext) {
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
