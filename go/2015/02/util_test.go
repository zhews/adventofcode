package main

import "testing"

func Test_parseInstruction(t *testing.T) {
	type args struct {
		instruction string
	}
	tests := []struct {
		name    string
		args    args
		want    int
		want1   int
		want2   int
		wantErr bool
	}{
		{
			name: "Example 1",
			args: args{
				instruction: "2x3x4",
			},
			want:    2,
			want1:   3,
			want2:   4,
			wantErr: false,
		},
		{
			name: "Example 2",
			args: args{
				instruction: "1x1x10",
			},
			want:    1,
			want1:   1,
			want2:   10,
			wantErr: false,
		},
		{
			name: "Invalid Instruction",
			args: args{
				instruction: "1x1",
			},
			want:    0,
			want1:   0,
			want2:   0,
			wantErr: true,
		},
		{
			name: "Invalid Length",
			args: args{
				instruction: "Ax1x2",
			},
			want:    0,
			want1:   0,
			want2:   0,
			wantErr: true,
		},
		{
			name: "Invalid Width",
			args: args{
				instruction: "1xAx2",
			},
			want:    0,
			want1:   0,
			want2:   0,
			wantErr: true,
		},
		{
			name: "Invalid Height",
			args: args{
				instruction: "1x2xA",
			},
			want:    0,
			want1:   0,
			want2:   0,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1, got2, err := parseInstruction(tt.args.instruction)
			if (err != nil) != tt.wantErr {
				t.Errorf("parseInstruction() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("parseInstruction() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("parseInstruction() got1 = %v, want %v", got1, tt.want1)
			}
			if got2 != tt.want2 {
				t.Errorf("parseInstruction() got2 = %v, want %v", got2, tt.want2)
			}
		})
	}
}
