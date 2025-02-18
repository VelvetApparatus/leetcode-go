package _079

import "testing"

func TestNumTilePossibilities(t *testing.T) {
	type args struct {
		tiles string
	}

	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Test 1",
			args: args{
				tiles: "AAB",
			},
			want: 8,
		},
		{
			name: "Test 2",
			args: args{
				tiles: "V",
			},
			want: 1,
		},
		{
			name: "Test 3",
			args: args{
				tiles: "AAABBC",
			},
			want: 188,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := numTilePossibilities(tt.args.tiles); got != tt.want {
				t.Errorf("numTilePossibilities() = %v, want %v", got, tt.want)
			}
		})
	}
}
