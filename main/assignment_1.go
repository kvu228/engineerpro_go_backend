package main

import "time"

type Person struct {
	BirthYear int
	Career    string
	Name      string
}

func (p *Person) CalculateAge() int {
	return time.Now().Year() - p.BirthYear
}

func (p *Person) CalculateJobMatch() bool {
	n := len(p.Name)
	if p.BirthYear%n == 0 {
		return true
	}
	return false
}
