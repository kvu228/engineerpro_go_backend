package util

import "fmt"

// InputSlice read inputs and return a slice of int
func InputSlice() []int {
	var n int
	fmt.Println("Enter n: ")
	_, err := fmt.Scan(&n)
	if err != nil {
		return nil
	}

	s := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Println("Enter element: ")
		_, err := fmt.Scan(&s[i])
		if err != nil {
			return nil
		}
	}
	return s
}
