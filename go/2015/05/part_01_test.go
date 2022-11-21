package main

import "testing"

func TestPartO1(t *testing.T) {
	type args struct {
		input string
	}
	tests := []struct {
		name    string
		args    args
		want    int
		wantErr bool
	}{
		{
			name: "Example 1",
			args: args{
				input: "ugknbfddgicrmopn",
			},
			want:    1,
			wantErr: false,
		},
		{
			name: "Example 2",
			args: args{
				input: "aaa",
			},
			want:    1,
			wantErr: false,
		},
		{
			name: "Example 3",
			args: args{
				input: "jchzalrnumimnmhp",
			},
			want:    0,
			wantErr: false,
		},
		{
			name: "Example 4",
			args: args{
				input: "haegwjzuvuyypxyu",
			},
			want:    0,
			wantErr: false,
		},
		{
			name: "Example 5",
			args: args{
				input: "dvszwmarrgswjxmb",
			},
			want:    0,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Part01(tt.args.input); got != tt.want {
				t.Errorf("Part01() = %v, want %v", got, tt.want)
			}
		})
	}
}
