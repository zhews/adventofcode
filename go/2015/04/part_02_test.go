package main

import "testing"

func TestPart02(t *testing.T) {
	type args struct {
		secretKey string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Example 1",
			args: args{
				secretKey: "abcdef",
			},
			want: 6742839,
		},
		{
			name: "Example 2",
			args: args{
				secretKey: "pqrstuv",
			},
			want: 5714438,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Part02(tt.args.secretKey); got != tt.want {
				t.Errorf("Part02() = %v, want %v", got, tt.want)
			}
		})
	}
}
