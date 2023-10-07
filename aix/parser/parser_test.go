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
			program, err := ParseFileByPath(tt.args.path)
			if !reflect.DeepEqual(err, tt.want) {
				t.Errorf("ParseFileByPath() = %v, want %v", err, tt.want)
			}
			println(program)
		})
	}
}
