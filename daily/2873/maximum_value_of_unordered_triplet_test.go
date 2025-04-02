package _2873

import "testing"

func Test_maximumTripletValue(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want int64
	}{
		{
			name: "case 1",
			args: args{
				nums: []int{12, 6, 1, 2, 7},
			},
			want: 77,
		},
		{
			name: "case 2",
			args: args{
				nums: []int{1, 10, 3, 4, 19},
			},
			want: 133,
		},
		{
			name: "case 3",
			args: args{
				nums: []int{1, 2, 3},
			},
			want: 0,
		},
		{
			name: "case 4",
			args: args{
				nums: []int{1000000, 1, 1000000},
			},
			want: 999999000000,
		},
		{
			name: "case 5",
			args: args{
				nums: []int{6, 11, 12, 12, 7, 9, 2, 11, 12, 4, 19, 14, 16, 8, 16},
			},
			want: 190,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maximumTripletValue(tt.args.nums); got != tt.want {
				t.Errorf("maximumTripletValue() = %v, want %v", got, tt.want)
			}
		})
	}
}
