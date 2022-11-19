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
				input: "2x3x4",
			},
			want:    58,
			wantErr: false,
		},
		{
			name: "Example 2",
			args: args{
				input: "1x1x10",
			},
			want:    43,
			wantErr: false,
		},
		{
			name: "Invalid Instruction",
			args: args{
				input: "1x2",
			},
			want:    0,
			wantErr: true,
		},
		{
			name: "Invalid Dimension",
			args: args{
				input: "Ax1x2",
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

func Test_calculatePaperAmountForPackage(t *testing.T) {
	type args struct {
		length int
		width  int
		height int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Example 1",
			args: args{
				length: 2,
				width:  3,
				height: 4,
			},
			want: 58,
		},
		{
			name: "Example 2",
			args: args{
				length: 1,
				width:  1,
				height: 10,
			},
			want: 43,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := calculatePaperAmountForPresent(tt.args.length, tt.args.width, tt.args.height); got != tt.want {
				t.Errorf("calculatePaperAmountForPresent() = %v, want %v", got, tt.want)
			}
		})
	}
}
