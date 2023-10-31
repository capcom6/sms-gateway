package services

import (
	"reflect"
	"testing"

	"github.com/capcom6/sms-gateway/internal/sms-gateway/models"
	"github.com/capcom6/sms-gateway/pkg/smsgateway"
)

func TestMessagesService_recipientsStateToModel(t *testing.T) {
	type args struct {
		input []smsgateway.RecipientState
	}
	tests := []struct {
		name string
		s    *MessagesService
		args args
		want []models.MessageRecipient
	}{
		{
			name: "Without +",
			s:    &MessagesService{},
			args: args{
				input: []smsgateway.RecipientState{
					{
						PhoneNumber: "79990001234",
						State:       "",
					},
				},
			},
			want: []models.MessageRecipient{
				{
					MessageID:   0,
					PhoneNumber: "+79990001234",
					State:       "",
				},
			},
		},
		{
			name: "With +",
			s:    &MessagesService{},
			args: args{
				input: []smsgateway.RecipientState{
					{
						PhoneNumber: "+79990001234",
						State:       "",
					},
				},
			},
			want: []models.MessageRecipient{
				{
					MessageID:   0,
					PhoneNumber: "+79990001234",
					State:       "",
				},
			},
		},
		{
			name: "Empty phone",
			s:    &MessagesService{},
			args: args{
				input: []smsgateway.RecipientState{
					{
						PhoneNumber: "",
						State:       "",
					},
				},
			},
			want: []models.MessageRecipient{
				{
					MessageID:   0,
					PhoneNumber: "",
					State:       "",
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.s.recipientsStateToModel(tt.args.input); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("MessagesService.recipientsStateToModel() = %v, want %v", got, tt.want)
			}
		})
	}
}
