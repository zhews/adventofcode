package main

import "testing"

func TestPart2(t *testing.T) {
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
				input: ")",
			},
			want:    1,
			wantErr: false,
		},
		{
			name: "Example 2",
			args: args{
				input: "()())",
			},
			want:    5,
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
		{
			name: "Never Entering the Basement",
			args: args{
				input: "(",
			},
			want:    0,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Part2(tt.args.input)
			if (err != nil) != tt.wantErr {
				t.Errorf("Part2() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Part2() got = %v, want %v", got, tt.want)
			}
		})
	}
}
