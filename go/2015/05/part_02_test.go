package main

import "testing"

func TestPart02(t *testing.T) {
	type args struct {
		input string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Example 1",
			args: args{
				input: "qjhvhtzxzqqjkmpb",
			},
			want: 1,
		},
		{
			name: "Example 2",
			args: args{
				input: "xxyxx",
			},
			want: 1,
		},
		{
			name: "Example 3",
			args: args{
				input: "uurcxstgmygtbstg",
			},
			want: 0,
		},
		{
			name: "Example 4",
			args: args{
				input: "ieodomkazucvgmuy",
			},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Part02(tt.args.input); got != tt.want {
				t.Errorf("Part02() = %v, want %v", got, tt.want)
			}
		})
	}
}
