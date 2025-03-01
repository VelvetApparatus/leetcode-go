package _460

import (
	"reflect"
	"testing"
)

func Test_applyOperations(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "case 1",
			args: args{
				nums: []int{1, 2, 2, 1, 1, 0},
			},
			want: []int{1, 4, 2, 0, 0, 0},
		},
		{
			name: "case 2",
			args: args{
				nums: []int{0, 1},
			},
			want: []int{1, 0},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := applyOperations(tt.args.nums); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("applyOperations() = %v, want %v", got, tt.want)
			}
		})
	}
}
