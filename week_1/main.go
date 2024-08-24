package main

import (
	"bufio"
	"engineerpro_go_backend/week_1/assignment_four"
	"engineerpro_go_backend/week_1/assignment_one"
	"engineerpro_go_backend/week_1/assignment_three"
	"engineerpro_go_backend/week_1/assignment_two"
	"fmt"
	"os"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Println("Choose assignment: ")
		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)
		switch input {
		case "1":
			assignment_one.AssignmentOneHandler()
		case "2":
			assignment_two.AssignmentTwoHandler()
		case "3":
			assignment_three.AssignmentThreeHandler()
		case "4":
			assignment_four.AssignmentFourHandler()
		}
	}

}
