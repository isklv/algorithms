package main

import (
	"testing"
)

func Test_replaceTwo(t *testing.T) {
	type args struct {
		a int
		b int
	}
	tests := []struct {
		name  string
		args  args
		want  int
		want1 int
	}{
		{
			name: "Result",
			args: args{
				a: 4,
				b: 8,
			},
			want:  8,
			want1: 4,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := swapTwo(tt.args.a, tt.args.b)
			if got != tt.want {
				t.Errorf("swapTwo() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("swapTwo() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}
