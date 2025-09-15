package bot

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/purkhanov/gogram/api"
)

const (
	baseUrl            = "https://api.telegram.org"
	contentTypeJSON    = "application/json"
	httpRequestTimeout = 5 * time.Second
)

type Bot struct {
	token        string
	urlWithToken string
	api          *api.ApiClient
	Ctx          context.Context
	Cancel       context.CancelFunc

	WebhookOptions *WebhookOptions
}

func NewBot(token string, webhookOptions *WebhookOptions) *Bot {
	ctx, cancel := context.WithCancel(context.Background())

	return &Bot{
		token:          token,
		urlWithToken:   fmt.Sprintf("%s/bot%s", baseUrl, token),
		api:            api.NewClient(&http.Client{}, ctx),
		Ctx:            ctx,
		Cancel:         cancel,
		WebhookOptions: webhookOptions,
	}
}
