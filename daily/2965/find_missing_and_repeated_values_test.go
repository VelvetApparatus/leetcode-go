package _2965

import (
	"reflect"
	"testing"
)

func Test_findMissingAndRepeatedValues(t *testing.T) {
	type args struct {
		grid [][]int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "case 1",
			args: args{
				grid: [][]int{{1, 3}, {2, 2}},
			},
			want: []int{2, 4},
		},
		{
			name: "case 2",
			args: args{
				grid: [][]int{{9, 1, 7}, {8, 9, 2}, {3, 4, 6}},
			},
			want: []int{9, 5},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findMissingAndRepeatedValues(tt.args.grid); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("findMissingAndRepeatedValues() = %v, want %v", got, tt.want)
			}
		})
	}
}
