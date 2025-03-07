package _2523

import (
	"reflect"
	"testing"
)

func Test_closestPrimes(t *testing.T) {
	type args struct {
		left  int
		right int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "case 1",
			args: args{
				left:  10,
				right: 19,
			},
			want: []int{11, 13},
		},
		{
			name: "case 2",
			args: args{
				left:  4,
				right: 6,
			},
			want: []int{-1, -1},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := closestPrimes(tt.args.left, tt.args.right); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("closestPrimes() = %v, want %v", got, tt.want)
			}
		})
	}
}
