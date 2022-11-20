package main

import "testing"

func TestPart01(t *testing.T) {
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
			want: 609043,
		},
		{
			name: "Example 2",
			args: args{
				secretKey: "pqrstuv",
			},
			want: 1048970,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Part01(tt.args.secretKey); got != tt.want {
				t.Errorf("Part01() = %v, want %v", got, tt.want)
			}
		})
	}
}
