package stringutil

import "testing"

func TestCommonStr(t *testing.T) {
	type args struct {
		a string
		b string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{name: "EmptyString", args: args{a: "Hello", b: "HelloWorld"}, want: "Hello"},
		{name: "NullString", args: args{a: "", b: ""}, want: ""},
		{name: "HellString", args: args{a: "Hell", b: "Hello"}, want: "Hell"},
		{name: "FirstEmptyString", args: args{a: "Hell", b: ""}, want: ""},
		{name: "SecondEmptyString", args: args{a: "", b: "Hello"}, want: ""},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CommonStr(tt.args.a, tt.args.b); got != tt.want {
				t.Errorf("CommonStr() = %v, want %v", got, tt.want)
			}
		})
	}
}
