package misc

import "testing"

func TestGetString(t *testing.T) {
	type args struct {
		m   map[string]interface{}
		key string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "key exists",
			args: args{
				m: map[string]interface{}{
					"key": "value",
				},
				key: "key",
			},
			want: "value",
		},
		{
			name: "key not found",
			args: args{
				m:   map[string]interface{}{},
				key: "key",
			},
			want: "",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetString(tt.args.m, tt.args.key); got != tt.want {
				t.Errorf("GetString() = %v, want %v", got, tt.want)
			}
		})
	}
}
