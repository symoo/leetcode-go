package twosum

import (
	"reflect"
	"testing"
)

func Test_brute(t *testing.T) {
	type args struct {
		nums   []int
		target int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{"sum", args {[]int{1, 2, 3}, 4,}, []int{0, 2},},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := brute(tt.args.nums, tt.args.target); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("brute() = %v, want %v", got, tt.want)
			}
		})
	}
}
