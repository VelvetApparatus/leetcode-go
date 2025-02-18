package _718

import (
	"reflect"
	"testing"
)

func Test_constructDistancedSequence(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "case 1",
			args: args{n: 3},
			want: []int{3, 1, 2, 3, 2},
		},
		{
			name: "case 2",
			args: args{n: 5},
			want: []int{5, 3, 1, 4, 3, 5, 2, 4, 2},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := constructDistancedSequence(tt.args.n); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("constructDistancedSequence() = %v, want %v", got, tt.want)
			}
		})
	}
}
