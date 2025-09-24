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
	httpRequestTimeout = 5 * time.Second
)

type Bot struct {
	urlWithToken string
	api          *api.ApiClient
	ctx          context.Context
}

func NewBot(ctx context.Context, token string) *Bot {
	return &Bot{
		urlWithToken: fmt.Sprintf("%s/bot%s", baseUrl, token),
		api:          api.NewClient(ctx, &http.Client{}),
		ctx:          ctx,
	}
}
