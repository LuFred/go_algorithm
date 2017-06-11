package tests

import (
	"go_algorithm/code/easy"
	"reflect"
	"testing"
)

func Test_Easy_Reverse(t *testing.T) {
	ti := map[int]int{
		123:          321,
		-123:         -321,
		123456789019: 0,
	}
	for k, v := range ti {
		if r := easy.Reverse(k); r != v {
			t.Error(
				"For",
				k,
				",Expected ",
				v,
				", got ",
				r)
		}
	}
}

func Test_Easy_TwoSum(t *testing.T) {
	nums := []int{3, 2, 4}
	targets := 6
	ans := []int{1, 2}
	r := easy.TwoSum(nums, targets)
	if !reflect.DeepEqual(r, ans) {
		t.Error(
			"Expected ",
			ans,
			" got ",
			r,
		)
	}
}
