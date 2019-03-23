package twosum

import (
	"reflect"
	"testing"
)

func Test_onePassHash(t *testing.T) {
	type args struct {
		nums   []int
		target int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{"two sum", args{[]int{7, 9, 8}, 17}, []int{1, 2}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := onePassHash(tt.args.nums, tt.args.target); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("onePassHash() = %v, want %v", got, tt.want)
			}
		})
	}
}
