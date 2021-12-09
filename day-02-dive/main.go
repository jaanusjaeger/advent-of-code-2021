package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	fname := os.Args[1]
	fmt.Println("Arguments:", os.Args)

	result1 := puzzle1(fname)
	result2 := puzzle2(fname)

	fmt.Println("RESULT1:", result1)
	fmt.Println("RESULT2:", result2)
}

func puzzle1(fname string) int {
	file, err := os.Open(fname)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	pos := 0
	depth := 0

	scanner := bufio.NewScanner(file)
	// optionally, resize scanner's capacity for lines over 64K, see next example
	for scanner.Scan() {
		line := scanner.Text()
		fmt.Println(line)
		if len(line) == 0 {
			continue
		}

		splits := strings.Split(line, " ")
		if len(splits) != 2 {
			log.Fatal("Invalid input", line)
		}
		cmd := splits[0]
		value, err := strconv.Atoi(splits[1])
		if err != nil {
			log.Fatal(err)
		}
		switch cmd {
		case "forward":
			pos += value
		case "down":
			depth += value
		case "up":
			depth += -1 * value
		}
		fmt.Println("pos =", pos, "depth =", depth)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return pos * depth
}

func puzzle2(fname string) int {
	file, err := os.Open(fname)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	aim := 0
	pos := 0
	depth := 0

	scanner := bufio.NewScanner(file)
	// optionally, resize scanner's capacity for lines over 64K, see next example
	for scanner.Scan() {
		line := scanner.Text()
		fmt.Println(line)
		if len(line) == 0 {
			continue
		}

		splits := strings.Split(line, " ")
		if len(splits) != 2 {
			log.Fatal("Invalid input", line)
		}
		cmd := splits[0]
		value, err := strconv.Atoi(splits[1])
		if err != nil {
			log.Fatal(err)
		}
		switch cmd {
		case "forward":
			pos += value
			depth += aim * value
		case "down":
			aim += value
		case "up":
			aim += -1 * value
		}

		fmt.Println("aim =", aim, "pos =", pos, "depth =", depth)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return pos * depth
}
