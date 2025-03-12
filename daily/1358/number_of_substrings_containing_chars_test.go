package _1358

import "testing"

func Test_numberOfSubstrings(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "case 1",
			args: args{s: "abcabc"},
			want: 10,
		},
		{
			name: "case 2",
			args: args{s: "aaacb"},
			want: 3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := numberOfSubstrings(tt.args.s); got != tt.want {
				t.Errorf("numberOfSubstrings() = %v, want %v", got, tt.want)
			}
		})
	}
}
