package main

import "testing"

func Test_binarySearch(t *testing.T) {
	type args struct {
		data   []int
		target int
		low    int
		high   int
	}
	tests := []struct {
		name      string
		args      args
		wantIndex int
		wantFound bool
	}{
		{
			name: "Search",
			args: args{
				data:   []int{1, 45, 76, 87, 94, 109, 555, 1003},
				target: 555,
				low:    0,
				high:   8,
			},
			wantIndex: 6,
			wantFound: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotIndex, gotFound := binarySearch(tt.args.data, tt.args.target, tt.args.low, tt.args.high)
			if gotIndex != tt.wantIndex {
				t.Errorf("binarySearch() gotIndex = %v, want %v", gotIndex, tt.wantIndex)
			}
			if gotFound != tt.wantFound {
				t.Errorf("binarySearch() gotFound = %v, want %v", gotFound, tt.wantFound)
			}
		})
	}
}
