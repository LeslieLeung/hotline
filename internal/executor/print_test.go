package executor

import (
	"github.com/leslieleung/hotline/internal/misc"
	"testing"
)

func TestPrint_BindParams(t *testing.T) {
	type fields struct {
		Message string
	}
	type args struct {
		params map[string]interface{}
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name:   "Valid parameters",
			fields: fields{Message: ""},
			args: args{
				params: map[string]interface{}{"message": "Hello, World!"},
			},
			wantErr: false,
		},
		{
			name:   "Missing parameters",
			fields: fields{Message: ""},
			args: args{
				params: map[string]interface{}{},
			},
			wantErr: false,
		},
		{
			name:   "Invalid parameter type",
			fields: fields{Message: ""},
			args: args{
				params: map[string]interface{}{"message": 12345},
			},
			wantErr: false, // Current implementation does not error on type mismatch, but this might not be the desired behavior
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &Print{
				Message: tt.fields.Message,
			}
			err := c.BindParams(tt.args.params)
			if (err != nil) != tt.wantErr {
				t.Errorf("BindParams() error = %v, wantErr %v", err, tt.wantErr)
			}
			// For valid parameters, also check if the message was correctly assigned
			if !tt.wantErr && c.Message != misc.GetString(tt.args.params, "message") {
				t.Errorf("BindParams() got = %v, want %v", c.Message, misc.GetString(tt.args.params, "message"))
			}
		})
	}
}
