package smsutility

import (
	"testing"
)

func TestSendSMS(t *testing.T) {
	type args struct {
		phoneNumber string
		message     string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "happy: sent message",
			args: args{
				phoneNumber: "0708113456",
				message:     "hello test",
			},
			wantErr: false,
		},
		{
			name: "sad: invalid phone number",
			args: args{
				phoneNumber: "070811",
				message:     "hello test",
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := SendSMS(tt.args.phoneNumber, tt.args.message); (err != nil) != tt.wantErr {
				t.Errorf("SendSMS() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestReplySMS(t *testing.T) {
	type args struct {
		phoneNumber string
		message     string
		linkID      string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "sad: invalid linkID",
			args: args{
				phoneNumber: "0708113456",
				message:     "hello test",
				linkID:      "123",
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := ReplySMS(tt.args.phoneNumber, tt.args.message, tt.args.linkID); (err != nil) != tt.wantErr {
				t.Errorf("ReplySMS() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
