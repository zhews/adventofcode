package main

import "testing"

func TestPart2(t *testing.T) {
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
				instructions: ")",
			},
			want:    1,
			wantErr: false,
		},
		{
			name: "Example 2",
			args: args{
				instructions: "()())",
			},
			want:    5,
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
		{
			name: "Never Entering the Basement",
			args: args{
				instructions: "(",
			},
			want:    0,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Part02(tt.args.instructions)
			if (err != nil) != tt.wantErr {
				t.Errorf("Part02() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Part02() got = %v, want %v", got, tt.want)
			}
		})
	}
}
