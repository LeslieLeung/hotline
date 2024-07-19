package executor

import (
	"reflect"
	"testing"
)

func TestCmd_Execute(t *testing.T) {
	type args struct {
		params map[string]interface{}
	}
	tests := []struct {
		name    string
		args    args
		want    map[string]interface{}
		wantErr bool
	}{
		{
			name: "echo",
			args: args{
				params: map[string]interface{}{
					"command": "echo \"hello, world\"",
				},
			},
			want: map[string]interface{}{
				"stdout": "hello, world\n",
				"items":  []string{},
			},
			wantErr: false,
		},
		{
			name: "echo with split",
			args: args{
				params: map[string]interface{}{
					"command":      `echo "alpha,beta,charlie"`,
					"split_output": ",",
				},
			},
			want: map[string]interface{}{
				"stdout": "alpha,beta,charlie\n",
				"items":  []string{"alpha", "beta", "charlie\n"},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &Cmd{}
			// Bind the parameters to the Cmd instance before execution
			err := c.BindParams(tt.args.params)
			if err != nil {
				t.Errorf("BindParams() error = %v", err)
				return
			}
			got, err := c.Execute()
			if (err != nil) != tt.wantErr {
				t.Errorf("Execute() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Execute() got = %v, want %v", got, tt.want)
			}
		})
	}
}
