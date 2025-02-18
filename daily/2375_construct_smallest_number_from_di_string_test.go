package daily

import "testing"

func Test_smallestNumber(t *testing.T) {
	type args struct {
		pattern string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "case1",
			args: args{"IIIDIDDD"},
			want: "123549876",
		},
		{
			name: "case2",
			args: args{"DDD"},
			want: "4321",
		},
		{
			name: "case3",
			args: args{"DDDIII"},
			want: "4321567",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := smallestNumber(tt.args.pattern); got != tt.want {
				t.Errorf("smallestNumber() = %v, want %v", got, tt.want)
			}
		})
	}
}
