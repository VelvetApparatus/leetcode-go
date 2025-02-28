package _092

import "testing"

func Test_shortestCommonSupersequence(t *testing.T) {
	type args struct {
		str1 string
		str2 string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "case 1",
			args: args{
				str1: "abac",
				str2: "cab",
			},
			want: "cabac",
		},
		{
			name: "case 2",
			args: args{
				str1: "aaaaaaaa",
				str2: "aaaaaaaa",
			},
			want: "aaaaaaaa",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := shortestCommonSupersequence(tt.args.str1, tt.args.str2); got != tt.want {
				t.Errorf("shortestCommonSupersequence() = %v, want %v", got, tt.want)
			}
		})
	}
}
