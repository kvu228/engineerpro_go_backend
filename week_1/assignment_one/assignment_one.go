package assignment_one

/*
Viết hàm nhập 2 cạnh của hình chữ nhật, in ra diện tích, chu vi
*/

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// readRectangleSides read the inputs and parse to float64
func readRectangleSides() (firstSide, secondSide float64) {
	reader := bufio.NewReader(os.Stdin)
	var err error
	for {
		fmt.Print("Enter first side: ")
		input, _ := reader.ReadString('\n')
		firstSide, err = strconv.ParseFloat(strings.TrimSpace(input), 64)
		if err != nil {
			fmt.Println("Error parsing first side", err)
			continue
		}
		if firstSide <= 0 {
			fmt.Println("first side is invalid, must greater than zero")
			continue
		}
		break
	}

	for {
		fmt.Print("Enter second side: ")
		input, _ := reader.ReadString('\n')
		secondSide, err = strconv.ParseFloat(strings.TrimSpace(input), 64)
		if err != nil {
			fmt.Println("Error parsing second side", err)
			continue
		}
		if secondSide <= 0 {
			fmt.Println("second side is invalid, must greater than zero")
			continue
		}
		break
	}

	return
}

// calculateRectangle receives sides and return the corresponding perimeter and area
func calculateRectangle(firstSide, secondSide float64) (perimeter float64, area float64) {
	perimeter = (firstSide + secondSide) * 2
	area = firstSide * secondSide
	return
}

func AssignmentOneHandler() {
	firstSide, secondSide := readRectangleSides()
	perimeter, area := calculateRectangle(firstSide, secondSide)
	fmt.Printf("Perimeter: %f, Area: %f\n", perimeter, area)
}
