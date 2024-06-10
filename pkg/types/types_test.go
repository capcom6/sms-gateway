package types

import (
	"testing"
)

func TestZeroDefault(t *testing.T) {
	tests := []struct {
		name  string
		value string
		def   string
		want  string
	}{
		{
			name:  "String zero value",
			value: "",
			def:   "default",
			want:  "default",
		},
		{
			name:  "String non-zero value",
			value: "value",
			def:   "default",
			want:  "value",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := ZeroDefault(tt.value, tt.def)
			if got != tt.want {
				t.Errorf("ZeroDefault() = %v, want %v", got, tt.want)
			}
		})
	}
}
