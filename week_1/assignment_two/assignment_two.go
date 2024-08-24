package assignment_two

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

/*
Viết chương trình nhập 1 string, in ra true nếu độ dài chuỗi chia hết cho 2, và false nếu ngược lại
*/

func AssignmentTwoHandler() bool {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter sentence: ")
	input, _ := reader.ReadString('\n')
	input = strings.TrimSuffix(input, "\n")
	if len(input)%2 == 0 {
		fmt.Println("ans: true")
		return true
	}
	fmt.Println("ans: false")
	return false
}
