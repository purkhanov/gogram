package bot

import (
	"errors"
	"fmt"
	"net/http"
	"net/url"

	"github.com/purkhanov/gogram/api"
	"github.com/purkhanov/gogram/types"
)

const (
	baseUrl         = "https://api.telegram.org"
	getUpdatesUrl   = "/getUpdates"
	contentTypeJSON = "application/json"
)

var errInvalidType = errors.New("invalid type")

type Bot struct {
	token        string
	urlWithToken string
	api          *api.ApiClient
}

func NewBot(token string) *Bot {
	return &Bot{
		token:        token,
		urlWithToken: fmt.Sprintf("%s/bot%s", baseUrl, token),
		api:          api.NewClient(&http.Client{}),
	}
}

func (b *Bot) GetUpdates(params url.Values) ([]types.Update, error) {
	req, err := http.NewRequest(http.MethodGet, b.urlWithToken+getUpdatesUrl, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %v", err)
	}
	req.URL.RawQuery = params.Encode()
	req.Header.Set("Content-Type", contentTypeJSON)

	resp, err := b.api.DoRequest(req)
	if err != nil {
		return nil, err
	}

	result, ok := resp.Result.([]types.Update)
	if !ok {
		return nil, errInvalidType
	}

	return result, nil
}
