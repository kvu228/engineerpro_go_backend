package assignment_four

import (
	"bufio"
	"engineerpro_go_backend/week_1/util"
	"fmt"
	"os"
	"strconv"
	"strings"
)

/*
Viết chương trình nhập giải bài toán twosum : https://leetcode.com/problems/two-sum/

Given an array of integers nums and an integer target, return indices of the two numbers such that they add up to target.

You may assume that each input would have exactly one solution, and you may not use the same element twice.

You can return the answer in any order.
*/

func twoSum(nums []int, target int) []int {
	myMap := make(map[int]int)
	for i, v := range nums {
		myMap[v] = i
	}

	for i, v := range nums {
		remainder := target - v
		if v, ok := myMap[remainder]; ok {
			if v != i {
				return []int{i, v}
			}
		}
	}
	return []int{-1, -1}
}

func AssignmentFourHandler() {
	s := util.InputSlice()
	var target int
	for {
		var err error
		fmt.Println("Input target: ")
		reader := bufio.NewReader(os.Stdin)
		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)
		target, err = strconv.Atoi(input)
		if err != nil {
			fmt.Println("input is not a number")
			continue
		}
		break
	}
	ans := twoSum(s, target)
	fmt.Printf("nums: %v\ttarget: %v\n", s, target)
	fmt.Println("ans: ", ans)
}
