package main

import (
	"bufio"
	"io"
	"os"
	"strconv"
	"strings"
)

func ProcessFile(filename string) ([]*Person, error) {
	var people []*Person

	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			return
		}
	}(file)

	reader := bufio.NewReader(file)
	for {
		line, err := reader.ReadString('\n')
		if err != nil {
			if err == io.EOF {
				break
			}
			return nil, err
		}
		line = strings.Trim(line, "\r\n")
		infos := strings.Split(line, "|")
		birthYear, err := strconv.ParseInt(infos[2], 10, 64)
		if err != nil {
			return nil, err
		}
		person := &Person{
			Name:      strings.ToUpper(infos[0]),
			Career:    strings.ToLower(infos[1]),
			BirthYear: int(birthYear),
		}
		people = append(people, person)
	}

	return people, nil
}
