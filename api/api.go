package api

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"net/http"
)

const contentTypeJSON = "application/json"

type ApiClient struct {
	httpClient *http.Client
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

func (c *ApiClient) DoRequestWithData(ctx context.Context, method, url string, data []byte) ([]byte, error) {
	req, err := http.NewRequestWithContext(ctx, method, url, bytes.NewBuffer(data))
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %w", err)
	}
	req.Header.Add("Content-Type", contentTypeJSON)

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("HTTP request failed: %w", err)
	}
	defer resp.Body.Close()

	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response body: %w", err)
	}

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("error status code: %s, Body: %s", resp.Status, string(respBody))
	}

	return respBody, nil
}
