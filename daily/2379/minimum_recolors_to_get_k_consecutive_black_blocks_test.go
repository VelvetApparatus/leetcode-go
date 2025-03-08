package _2379

import "testing"

func Test_minimumRecolors(t *testing.T) {
	type args struct {
		blocks string
		k      int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "case 1",
			args: args{
				blocks: "WBBWWBBWBW",
				k:      7,
			},
			want: 3,
		},
		{
			name: "case 2",
			args: args{
				blocks: "WBWBBBW",
				k:      2,
			},
			want: 0,
		},
		{
			name: "case 3",
			args: args{
				blocks: "BWWWBB",
				k:      6,
			},
			want: 3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := minimumRecolors(tt.args.blocks, tt.args.k); got != tt.want {
				t.Errorf("minimumRecolors() = %v, want %v", got, tt.want)
			}
		})
	}
}
