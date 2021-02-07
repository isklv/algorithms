package main

import "testing"

func Test_radixSort(t *testing.T) {
	type args struct {
		array []int
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "Sort",
			args: args{array: generateRadomArray(100, 0, 100)},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			radixSort(tt.args.array)
			for i := 0; i < len(tt.args.array)-1; i++ {
				if tt.args.array[i] > tt.args.array[i+1] {
					t.Errorf("radixSort() doesn't have valid result = %v > %v", tt.args.array[i], tt.args.array[i+1])
				}
			}
		})
	}
}
