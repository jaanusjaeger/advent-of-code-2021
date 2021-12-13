package main

import (
	"fmt"
	"strings"
)

func main() {
	file := openFileFromArgs()
	defer file.Close()

	scanner := func(line string) error {
		ss := strings.Split(line, ",")
		if len(ss) != 2 {
			return fmt.Errorf("invalid input: %s", line)
		}
		// TODO
		return nil
	}
	scanLines(file, scanner)

	fmt.Println("INPUT:")
	//fmt.Println(input)

	result1 := puzzle1()
	result2 := puzzle2()

	fmt.Println("RESULT1:", result1)
	fmt.Println("RESULT2:", result2)
}

func puzzle1() string {
	result := ""

	return result
}

func puzzle2() string {
	result := ""

	return result
}
