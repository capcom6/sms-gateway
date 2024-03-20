package upstream

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"sync"

	"github.com/capcom6/sms-gateway/pkg/smsgateway"
)

const BASE_URL = "https://sms.capcom.me/api/upstream/v1"

type Client struct {
	options map[string]string

	client *http.Client
	mux    sync.Mutex
}

func New(options map[string]string) (*Client, error) {
	return &Client{
		options: options,
	}, nil
}

func (c *Client) Open(ctx context.Context) error {
	c.mux.Lock()
	defer c.mux.Unlock()

	if c.client != nil {
		return nil
	}

	c.client = &http.Client{}

	return nil
}

func (c *Client) Send(ctx context.Context, messages map[string]map[string]string) error {
	payload := make(smsgateway.UpstreamPushRequest, 0, len(messages))

	for address := range messages {
		payload = append(payload, smsgateway.PushNotification{
			Token: address,
		})
	}

	payloadBytes, err := json.Marshal(payload)
	if err != nil {
		return fmt.Errorf("can't marshal payload: %w", err)
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodPost, BASE_URL+"/push", bytes.NewReader(payloadBytes))
	if err != nil {
		return fmt.Errorf("can't create request: %w", err)
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("User-Agent", "android-sms-gateway/1.x (server; golang)")

	resp, err := c.client.Do(req)
	if err != nil {
		return fmt.Errorf("can't send request: %w", err)
	}

	defer func() {
		_, _ = io.Copy(io.Discard, resp.Body)
		_ = resp.Body.Close()
	}()

	if resp.StatusCode >= 400 {
		return fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}

	return nil
}

func (c *Client) Close(ctx context.Context) error {
	c.mux.Lock()
	defer c.mux.Unlock()

	c.client = nil

	return nil
}
