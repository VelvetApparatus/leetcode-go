package _415

import "testing"

func Test_getHappyString(t *testing.T) {
	type args struct {
		n int
		k int
	}

	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "case 1",
			args: args{
				n: 1, k: 3,
			},
			want: "c",
		},
		{
			name: "case 2",
			args: args{
				n: 1, k: 4,
			},
			want: "",
		},
		{
			name: "case 3",
			args: args{
				n: 3, k: 9,
			},
			want: "cab",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getHappyString(tt.args.n, tt.args.k); got != tt.want {
				t.Errorf("getHappyString() = %v, want %v", got, tt.want)
			}
		})
	}
}
