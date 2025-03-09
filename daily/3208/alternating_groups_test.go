package _3208

import "testing"

func Test_numberOfAlternatingGroups(t *testing.T) {
	type args struct {
		colors []int
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
				colors: []int{0, 1, 0, 1, 0},
				k:      3,
			},
			want: 3,
		},
		{
			name: "case 2",
			args: args{
				colors: []int{0, 1, 0, 0, 1, 0, 1},
				k:      6,
			},
			want: 2,
		},
		{
			name: "case 3",
			args: args{
				colors: []int{1, 1, 0, 1},
				k:      4,
			},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := numberOfAlternatingGroups(tt.args.colors, tt.args.k); got != tt.want {
				t.Errorf("numberOfAlternatingGroups() = %v, want %v", got, tt.want)
			}
		})
	}
}
