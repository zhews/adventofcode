package main

import "testing"

func Test_minInt(t *testing.T) {
	type args struct {
		digits []int
	}
	tests := []struct {
		name    string
		args    args
		want    int
		wantErr bool
	}{
		{
			name: "1 to 5",
			args: args{
				digits: []int{1, 2, 3, 4, 5},
			},
			want:    1,
			wantErr: false,
		},
		{
			name: "5 to 1",
			args: args{
				digits: []int{5, 4, 3, 2, 1},
			},
			want:    1,
			wantErr: false,
		},
		{
			name: "Empty Array",
			args: args{
				digits: []int{},
			},
			want:    0,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := minInt(tt.args.digits...)
			if (err != nil) != tt.wantErr {
				t.Errorf("minInt() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("minInt() got = %v, want %v", got, tt.want)
			}
		})
	}
}
