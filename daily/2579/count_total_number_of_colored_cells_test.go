package _2579

import "testing"

func Test_coloredCells(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want int64
	}{
		{
			name: "case 1",
			args: args{n: 1},
			want: 1,
		},
		{
			name: "case 2",
			args: args{n: 2},
			want: 5,
		},
		{
			name: "case 3",
			args: args{n: 3},
			want: 13,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := coloredCells(tt.args.n); got != tt.want {
				t.Errorf("coloredCells() = %v, want %v", got, tt.want)
			}
		})
	}
}
