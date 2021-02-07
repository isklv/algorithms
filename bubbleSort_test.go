package main

import (
	"testing"
)

func Test_bubbleSort(t *testing.T) {
	type args struct {
		data []int
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "Sort",
			args: args{data: generateRadomArray(100, 0, 100)},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			bubbleSort(tt.args.data)
			for i := 0; i < len(tt.args.data)-1; i++ {
				if tt.args.data[i] > tt.args.data[i+1] {
					t.Errorf("bubbleSort() doesn't have valid result = %v > %v", tt.args.data[i], tt.args.data[i+1])
				}
			}
		})
	}
}
