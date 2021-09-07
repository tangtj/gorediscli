package cli

import (
	"reflect"
	"testing"
)

func TestGenByte(t *testing.T) {
	type args struct {
		command []byte
	}
	var tests = []struct {
		name string
		args args
		want [][]byte
	}{
		{
			name: "set 1 1",
			args: args{command: []byte("set 1 1")},
			want: func() [][]byte {
				r := make([][]byte, 0)
				r = append(r, []byte("set"))
				r = append(r, []byte("1"))
				r = append(r, []byte("1"))
				return r
			}(),
		},
		{
			name: "set 1 '1'",
			args: args{command: []byte("set 1 '1'")},
			want: func() [][]byte {
				r := make([][]byte, 0)
				r = append(r, []byte("set"))
				r = append(r, []byte("1"))
				r = append(r, []byte("'1'"))
				return r
			}(),
		},
		{
			name: "set 1 \"aksjhd aksjhdahsdkahsdlh 2alksdnhl\"",
			args: args{command: []byte("set 1 \"aksjhd aksjhdahsdkahsdlh 2alksdnhl\"")},
			want: func() [][]byte {
				r := make([][]byte, 0)
				r = append(r, []byte("set"))
				r = append(r, []byte("1"))
				r = append(r, []byte("\"aksjhd aksjhdahsdkahsdlh 2alksdnhl\""))
				return r
			}(),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GenByte(tt.args.command); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GenRedisCommand() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestConvertToCommand(t *testing.T) {
	type args struct {
		bs []byte
	}
	tests := []struct {
		name string
		args args
		want []byte
	}{
		//{
		//	name: "set 1 1",
		//	args: args{bs: []byte("set 1 1")},
		//},
		{
			name: "set mykey myvalue",
			args: args{bs: []byte("set mykey myvalue")},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ConvertToCommand(tt.args.bs); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ConvertToCommand() = %v, want %v", got, tt.want)
			}
		})
	}
}
