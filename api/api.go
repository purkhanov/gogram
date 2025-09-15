package api

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"net/http"
	"time"
)

const (
	httpRequestTimeout = 5 * time.Second
	contentTypeJSON    = "application/json"
)

type ApiClient struct {
	httpClient *http.Client
	ctx        context.Context
}

func NewClient(httpClient *http.Client, ctx context.Context) *ApiClient {
	return &ApiClient{
		httpClient: httpClient,
		ctx:        ctx,
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

func (c *ApiClient) DoRequestWithData(method, url string, data []byte) ([]byte, error) {
	req, err := http.NewRequest(method, url, bytes.NewBuffer(data))
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %w", err)
	}
	req.Header.Add("Content-Type", contentTypeJSON)

	return c.DoRequest(req)
}

func (c *ApiClient) DoRequestWithContextAndData(ctx context.Context, method, url string, data []byte) ([]byte, error) {
	req, err := http.NewRequestWithContext(ctx, method, url, bytes.NewBuffer(data))
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %w", err)
	}
	req.Header.Add("Content-Type", contentTypeJSON)

	return c.DoRequest(req)
}

func (c *ApiClient) DoRequestWithTimeout(method, url string, data []byte) ([]byte, error) {
	ctx, cancel := context.WithTimeout(c.ctx, httpRequestTimeout)
	defer cancel()

	req, err := http.NewRequestWithContext(ctx, method, url, bytes.NewBuffer(data))
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %w", err)
	}
	req.Header.Add("Content-Type", contentTypeJSON)

	return c.DoRequest(req)
}
