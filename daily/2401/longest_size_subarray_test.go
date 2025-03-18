package _2401

import "testing"

func Test_longestNiceSubarray(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "case 1",
			args: args{
				nums: []int{1, 3, 8, 48, 10},
			},
			want: 3,
		},
		{
			name: "case 2",
			args: args{
				nums: []int{3, 1, 5, 11, 13},
			},
			want: 1,
		},
		{
			name: "case 3",
			args: args{
				nums: []int{135745088, 609245787, 16, 2048, 2097152},
			},
			want: 3,
		},
		{
			name: "case 4",
			args: args{
				nums: []int{84139415, 693324769, 614626365, 497710833, 615598711, 264, 65552, 50331652, 1, 1048576, 16384, 544, 270532608, 151813349, 221976871, 678178917, 845710321, 751376227, 331656525, 739558112, 267703680},
			},
			want: 8,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := longestNiceSubarray(tt.args.nums); got != tt.want {
				t.Errorf("longestNiceSubarray() = %v, want %v", got, tt.want)
			}
		})
	}
}
