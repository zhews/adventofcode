package main

import "testing"

func TestPart1(t *testing.T) {
	type args struct {
		instructions string
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
				instructions: "(())",
			},
			want:    0,
			wantErr: false,
		},
		{
			name: "Example 2",
			args: args{
				instructions: "()()",
			},
			want:    0,
			wantErr: false,
		},
		{
			name: "Example 3",
			args: args{
				instructions: "(((",
			},
			want:    3,
			wantErr: false,
		},
		{
			name: "Example 4",
			args: args{
				instructions: "(()(()(",
			},
			want:    3,
			wantErr: false,
		},
		{
			name: "Example 5",
			args: args{
				instructions: "))(((((",
			},
			want:    3,
			wantErr: false,
		},
		{
			name: "Example 6",
			args: args{
				instructions: "())",
			},
			want:    -1,
			wantErr: false,
		},
		{
			name: "Example 7",
			args: args{
				instructions: "))(",
			},
			want:    -1,
			wantErr: false,
		},
		{
			name: "Example 8",
			args: args{
				instructions: ")))",
			},
			want:    -3,
			wantErr: false,
		},
		{
			name: "Example 9",
			args: args{
				instructions: ")())())",
			},
			want:    -3,
			wantErr: false,
		},
		{
			name: "Empty Input",
			args: args{
				instructions: "",
			},
			want:    0,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Part01(tt.args.instructions)
			if (err != nil) != tt.wantErr {
				t.Errorf("Part01() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Part01() got = %v, want %v", got, tt.want)
			}
		})
	}
}
