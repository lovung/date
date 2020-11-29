package num

import "testing"

func TestSign(t *testing.T) {
	type args struct {
		num int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{args: args{-12}, want: -1},
		{args: args{-6}, want: -1},
		{args: args{0}, want: 0},
		{args: args{6}, want: 1},
		{args: args{12}, want: 1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Sign(tt.args.num); got != tt.want {
				t.Errorf("Sign() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMaxInt(t *testing.T) {
	type args struct {
		first int
		rest  []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{args: args{0, []int{-1, -2, -3, -4, -5}}, want: 0},
		{args: args{12, []int{-1, -2, -3, -4, -5}}, want: 12},
		{args: args{-12, []int{-1, -2, -3, -4, -5}}, want: -1},
		{args: args{0, []int{1, 2, 3, 4, 5}}, want: 5},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := MaxInt(tt.args.first, tt.args.rest...); got != tt.want {
				t.Errorf("MaxInt() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMinInt(t *testing.T) {
	type args struct {
		first int
		rest  []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{args: args{0, []int{-1, -2, -3, -4, -5}}, want: -5},
		{args: args{12, []int{-1, -2, -3, -4, -5}}, want: -5},
		{args: args{-12, []int{-1, -2, -3, -4, -5}}, want: -12},
		{args: args{0, []int{1, 2, 3, 4, 5}}, want: 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := MinInt(tt.args.first, tt.args.rest...); got != tt.want {
				t.Errorf("MinInt() = %v, want %v", got, tt.want)
			}
		})
	}
}
