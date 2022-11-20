package main

import "testing"

func Test_mine(t *testing.T) {
	type args struct {
		secretKey      string
		requiredPrefix string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Example 1",
			args: args{
				secretKey:      "abcdef",
				requiredPrefix: "00000",
			},
			want: 609043,
		},
		{
			name: "Example 2",
			args: args{
				secretKey:      "pqrstuv",
				requiredPrefix: "00000",
			},
			want: 1048970,
		},
		{
			name: "Example 3",
			args: args{
				secretKey:      "abcdef",
				requiredPrefix: "000000",
			},
			want: 6742839,
		},
		{
			name: "Example 4",
			args: args{
				secretKey:      "pqrstuv",
				requiredPrefix: "000000",
			},
			want: 5714438,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := mine(tt.args.secretKey, tt.args.requiredPrefix); got != tt.want {
				t.Errorf("mine() = %v, want %v", got, tt.want)
			}
		})
	}
}
