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
	ctx          context.Context
}

func NewBot(token string) *Bot {
	ctx := context.Background()

	return &Bot{
		token:        token,
		urlWithToken: fmt.Sprintf("%s/bot%s", baseUrl, token),
		api:          api.NewClient(&http.Client{}),
		ctx:          ctx,
	}
}
