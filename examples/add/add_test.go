package add

import (
	"testing"
)

func TestAdd(t *testing.T) {
	type args struct {
		a int
		b int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Addition of two positive numbers",
			args: args{a: 2, b: 3},
			want: 5,
		},
		{
			name: "Addition of two negative numbers",
			args: args{a: -2, b: -3},
			want: -5,
		},
		{
			name: "Addition of positive and negative numbers",
			args: args{a: 2, b: -3},
			want: -1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Add(tt.args.a, tt.args.b); got != tt.want {
				t.Errorf("Add() = %v, want %v", got, tt.want)
			}
		})
	}
}
