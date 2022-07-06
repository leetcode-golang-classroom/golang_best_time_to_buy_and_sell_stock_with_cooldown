package sol

import "testing"

func Benchmark(b *testing.B) {
	prices := []int{1, 2, 3, 0, 2}
	for idx := 0; idx < b.N; idx++ {
		maxProfit(prices)
	}
}
func Test_maxProfit(t *testing.T) {
	type args struct {
		prices []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "prices = [1,2,3,0,2]",
			args: args{prices: []int{1, 2, 3, 0, 2}},
			want: 3,
		},
		{
			name: "prices = [1]",
			args: args{prices: []int{1}},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxProfit(tt.args.prices); got != tt.want {
				t.Errorf("maxProfit() = %v, want %v", got, tt.want)
			}
		})
	}
}
