package _749

import "testing"

func Test_maxAbsoluteSum(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "case 1",
			args: args{
				nums: []int{1, -3, 2, 3, -4},
			},
			want: 5,
		},
		{
			name: "case 2",
			args: args{
				nums: []int{2, -5, 1, -4, 3, -2},
			},
			want: 8,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxAbsoluteSum(tt.args.nums); got != tt.want {
				t.Errorf("maxAbsoluteSum() = %v, want %v", got, tt.want)
			}
		})
	}
}
