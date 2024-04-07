package messages

import (
	"context"
	"testing"
)

func TestModel_setCurrency(t *testing.T) {
	type fields struct {
		tgClient MessageSender
		config   ConfigGetter
		userDB   UserDB
	}
	type args struct {
		ctx context.Context
		msg Message
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    string
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Model{
				tgClient: tt.fields.tgClient,
				config:   tt.fields.config,
				userDB:   tt.fields.userDB,
			}
			got, err := s.setCurrency(tt.args.ctx, tt.args.msg)
			if (err != nil) != tt.wantErr {
				t.Errorf("setCurrency() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("setCurrency() got = %v, want %v", got, tt.want)
			}
		})
	}
}