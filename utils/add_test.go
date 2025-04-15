package utils

import "testing"

func Test_add(t *testing.T) {
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
			name: "add",
			args: args{
				a: 1,
				b: 2,
			},
			want: 3,
		},
		{
			name: "add",
			args: args{
				a: 3,
				b: 4,
			},
			want: 7,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Add(tt.args.a, tt.args.b); got != tt.want {
				t.Errorf("add() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAdd1add2(t *testing.T) {
	if got := Add(1, 2); got != 3 {
		t.Errorf("Add() = %v, want %v", got, 3)
	}
}

func TestAdd3add4(t *testing.T) {
	if got := Add(3, 4); got != 7 {
		t.Errorf("Add() = %v, want %v", got, 3)
	}
}
