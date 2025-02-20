package _980

import "testing"

func Test_findDifferentBinaryString(t *testing.T) {
	type args struct {
		nums []string
	}
	tests := []struct {
		name    string
		args    args
		wantAny []string
	}{
		{
			name: "case 1",
			args: args{
				nums: []string{"01", "10"},
			},
			wantAny: []string{"00", "11"},
		},
		{
			name: "case 2",
			args: args{
				nums: []string{"111", "011", "001"},
			},
			wantAny: []string{"101", "000", "010", "100", "110"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := findDifferentBinaryString(tt.args.nums)
			correct := false
			for _, want := range tt.wantAny {
				if want == got {
					correct = true
				}
			}
			if !correct {
				t.Errorf("findDifferentBinaryString(): want %v, got %v", tt.wantAny, got)
			}
		})
	}
}
