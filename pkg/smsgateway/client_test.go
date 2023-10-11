package smsgateway

import (
	"context"
	"io"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"
)

func TestClient_Send(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/message" {
			w.WriteHeader(http.StatusNotFound)
			return
		}

		if r.Header.Get("Content-Type") != "application/json" {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		req, _ := io.ReadAll(r.Body)
		defer r.Body.Close()

		if string(req) != `{"message":"","phoneNumbers":null}` {
			w.WriteHeader(http.StatusBadRequest)
			w.Write(req)
			return
		}

		w.WriteHeader(http.StatusCreated)
		w.Write([]byte(`{}`))
	}))
	defer server.Close()

	client := NewClient(Config{
		BaseURL: server.URL,
	})

	type args struct {
		ctx     context.Context
		message Message
	}
	tests := []struct {
		name    string
		c       *Client
		args    args
		want    MessageState
		wantErr bool
	}{
		{
			name: "Success",
			c:    client,
			args: args{
				ctx:     context.TODO(),
				message: Message{},
			},
			want:    MessageState{},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.c.Send(tt.args.ctx, tt.args.message)
			if (err != nil) != tt.wantErr {
				t.Errorf("Client.Send() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Client.Send() = %v, want %v", got, tt.want)
			}
		})
	}
}
