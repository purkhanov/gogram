package api

import (
	"fmt"
	"io"
	"net/http"
)


type ApiClient struct {
	httpClient *http.Client
	bufferSize int
	logger     any
}

func NewClient(httpClient *http.Client) *ApiClient {
	return &ApiClient{
		httpClient: httpClient,
	}
}

func (c *ApiClient) DoRequest(request *http.Request) ([]byte, error) {
	response, err := c.httpClient.Do(request)
	if err != nil {
		return nil, fmt.Errorf("HTTP request failed: %w", err)
	}
	defer response.Body.Close()

	respBody, err := io.ReadAll(response.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response body: %w", err)
	}

	if response.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("error status code: %s, Body: %s", response.Status, string(respBody))
	}

	return respBody, nil
}
