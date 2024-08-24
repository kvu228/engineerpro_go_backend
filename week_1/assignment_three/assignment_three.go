package assignment_three

import (
	"engineerpro_go_backend/week_1/util"
	"fmt"
)

/*
Viết chương trình nhập một slice số, in ra tổng, số lớn nhất, số nhỏ nhất, trung bình cộng, slice đã được sắp xếp
*/

func sumSlice(s []int) (sum int) {
	for i := 0; i < len(s); i++ {
		sum += s[i]
	}
	return
}

func minSlice(s []int) (min int) {
	for i := 0; i < len(s); i++ {
		if s[i] < min {
			min = s[i]
		}
	}
	return
}

func maxSlice(s []int) (max int) {
	for i := 0; i < len(s); i++ {
		if s[i] > max {
			max = s[i]
		}
	}
	return
}

func sortSlice(s []int) (sorted []int) {
	return quickSort(s, 0, len(s)-1)
}

func partition(s []int, low, high int) ([]int, int) {
	pivot := s[high]
	i := low
	for j := low; j < high; j++ {
		if s[j] < pivot {
			s[i], s[j] = s[j], s[i]
			i++
		}
	}
	s[i], s[high] = s[high], s[i]
	return s, i
}

func quickSort(s []int, low, high int) []int {
	if low < high {
		var pivot int
		s, pivot = partition(s, low, high)
		s = quickSort(s, low, pivot-1)
		s = quickSort(s, pivot+1, high)
	}
	return s
}

func AssignmentThreeHandler() {
	s := util.InputSlice()
	sSum := sumSlice(s)
	sMin := minSlice(s)
	sMax := maxSlice(s)
	sortedSlice := sortSlice(s)
	fmt.Printf("sum = %v\nmin = %v\nmax = %v\n", sSum, sMin, sMax)
	fmt.Println("sorted: ", sortedSlice)
}
