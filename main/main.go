package main

import "fmt"

func main() {

	// Example assignment 4
	filename := "main/a.txt"
	people, err := ProcessFile(filename)
	if err != nil {
		panic(err)
	}
	for _, person := range people {
		fmt.Println(person)
	}

	// Example assignment 1
	for _, person := range people {
		fmt.Printf("Name: %s, Age: %d\n", person.Name, person.CalculateAge())
	}

	// Example assignment 2
	example := "hello engineer pro"
	myMap := CreateMap(example)
	fmt.Println(myMap)
}
