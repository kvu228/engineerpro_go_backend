package main

func TwoSum(nums []int, target int) []int {
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
