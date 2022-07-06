package main

import (
	"fmt"
	"reflect"
)

/**
两数之和
链接: https://leetcode.cn/problems/two-sum/
时间复杂度：O(n)
空间复杂度: O(n)
*/

func twoSum(nums []int, target int) []int {
	m := make(map[int]int, len(nums))
	for i, num := range nums {
		_i, ok := m[target-num]
		if ok {
			return []int{_i, i}
		} else {
			m[num] = i
		}
	}
	return []int{-1, -1}
}

func main() {
	type Input struct {
		Nums   []int
		Target int
	}
	type Case struct {
		Input  Input
		Expect []int
	}
	cases := []Case{
		{
			Input: Input{
				Nums:   []int{2, 7, 11, 15},
				Target: 9,
			},
			Expect: []int{0, 1},
		},
		{
			Input: Input{
				Nums:   []int{3, 2, 4},
				Target: 6,
			},
			Expect: []int{1, 2},
		},
		{
			Input:  Input{Nums: []int{3, 3}, Target: 6},
			Expect: []int{0, 1},
		},
	}

	for _, _case := range cases {
		output := twoSum(_case.Input.Nums, _case.Input.Target)
		if !reflect.DeepEqual(output, _case.Expect) {
			fmt.Printf("expect: %v. output: %v\n", _case.Expect, output)
		}
	}
}
