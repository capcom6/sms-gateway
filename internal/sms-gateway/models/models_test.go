package models_test

import (
	"testing"

	"github.com/capcom6/sms-gateway/internal/sms-gateway/models"
)

func TestDevice_IsEmpty(t *testing.T) {
	tests := []struct {
		name string
		d    *models.Device
		want bool
	}{
		{
			name: "nil Device",
			d:    nil,
			want: true,
		},
		{
			name: "empty ID",
			d: &models.Device{
				ID: "",
			},
			want: true,
		},
		{
			name: "non-empty ID",
			d: &models.Device{
				ID: "some-id",
			},
			want: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.d.IsEmpty(); got != tt.want {
				t.Errorf("IsEmpty() = %v, want %v", got, tt.want)
			}
		})
	}
}
