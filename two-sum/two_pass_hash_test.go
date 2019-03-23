package twosum

import (
	"reflect"
	"testing"
)

func Test_twoPass(t *testing.T) {
	type args struct {
		nums   []int
		target int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{"two sum", args{[]int{1, 2, 3}, 4}, []int{0, 2}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := twoPass(tt.args.nums, tt.args.target); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("twoPass() = %v, want %v", got, tt.want)
			}
		})
	}
}
