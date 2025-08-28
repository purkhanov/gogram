package bot

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/purkhanov/gogram/api"
)

const (
	baseUrl         = "https://api.telegram.org"
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