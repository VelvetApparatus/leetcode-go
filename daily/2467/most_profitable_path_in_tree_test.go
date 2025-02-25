package _467

import "testing"

func Test_mostProfitablePath(t *testing.T) {
	type args struct {
		edges  [][]int
		bob    int
		amount []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "case 1",
			args: args{
				edges:  [][]int{{0, 1}, {1, 2}, {1, 3}, {3, 4}},
				bob:    3,
				amount: []int{-2, 4, 2, -4, 6},
			},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := mostProfitablePath(tt.args.edges, tt.args.bob, tt.args.amount); got != tt.want {
				t.Errorf("mostProfitablePath() = %v, want %v", got, tt.want)
			}
		})
	}
}
