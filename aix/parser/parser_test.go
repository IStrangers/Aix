package parser

import (
	"reflect"
	"testing"
)

func TestParseFileByPath(t *testing.T) {
	type args struct {
		path string
	}
	tests := []struct {
		name string
		args args
		want any
	}{
		{
			name: "HelloAix",
			args: args{
				path: "../../example/HelloAix.aix",
			},
			want: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ParseFileByPath(tt.args.path); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ParseFileByPath() = %v, want %v", got, tt.want)
			}
		})
	}
}
