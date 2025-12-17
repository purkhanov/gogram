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

	"github.com/purkhanov/gogram/api"
	"github.com/purkhanov/gogram/bot"
	"github.com/purkhanov/gogram/types"
)

const webhookSecretToken = "X-Telegram-Bot-Api-Secret-Token"

func (d *Dispatcher) StartWebhookServer(port uint16, options bot.WebhookOptions) error {
	if port == 0 {
		return errors.New("port cannot be zero")
	}

	handler := func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
			return
		}

		if options.SecretToken != "" {
			receivedToken := r.Header.Get(webhookSecretToken)
			expectedToken := options.SecretToken

			if receivedToken != expectedToken {
				log.Printf("Invalid secret token: received '%s', expected '%s'", receivedToken, expectedToken)
				http.Error(w, "Invalid secret token", http.StatusUnauthorized)
				return
			}
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
		case d.updatesChan <- update:
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
		ctx.JSON(http.StatusBadRequest, api.ApiResponse{
			Message: "Invalid JSON",
			Error:   err.Error(),
		})

		return
	}

	go d.checkUpdate(update)

	ctx.JSON(http.StatusOK, api.ApiResponse{Status: "accepted"})

}
