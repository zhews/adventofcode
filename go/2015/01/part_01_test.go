package main

import "testing"

func TestPart1(t *testing.T) {
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
				input: "(())",
			},
			want:    0,
			wantErr: false,
		},
		{
			name: "Example 2",
			args: args{
				input: "()()",
			},
			want:    0,
			wantErr: false,
		},
		{
			name: "Example 3",
			args: args{
				input: "(((",
			},
			want:    3,
			wantErr: false,
		},
		{
			name: "Example 4",
			args: args{
				input: "(()(()(",
			},
			want:    3,
			wantErr: false,
		},
		{
			name: "Example 5",
			args: args{
				input: "))(((((",
			},
			want:    3,
			wantErr: false,
		},
		{
			name: "Example 6",
			args: args{
				input: "())",
			},
			want:    -1,
			wantErr: false,
		},
		{
			name: "Example 7",
			args: args{
				input: "))(",
			},
			want:    -1,
			wantErr: false,
		},
		{
			name: "Example 8",
			args: args{
				input: ")))",
			},
			want:    -3,
			wantErr: false,
		},
		{
			name: "Example 9",
			args: args{
				input: ")())())",
			},
			want:    -3,
			wantErr: false,
		},
		{
			name: "Empty Input",
			args: args{
				input: "",
			},
			want:    0,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Part1(tt.args.input)
			if (err != nil) != tt.wantErr {
				t.Errorf("Part1() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Part1() got = %v, want %v", got, tt.want)
			}
		})
	}
}
