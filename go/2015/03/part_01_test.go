package main

import "testing"

func TestPart1(t *testing.T) {
	type args struct {
		directions string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Example 1",
			args: args{
				directions: ">",
			},
			want: 2,
		},
		{
			name: "Example 2",
			args: args{
				directions: "^>v<",
			},
			want: 4,
		},
		{
			name: "Example 3",
			args: args{
				directions: "^v^v^v^v^v",
			},
			want: 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Part1(tt.args.directions); got != tt.want {
				t.Errorf("Part1() = %v, want %v", got, tt.want)
			}
		})
	}
}
