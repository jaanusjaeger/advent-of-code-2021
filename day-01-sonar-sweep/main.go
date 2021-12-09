package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
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

	prev := -1
	count := 0
	scanner := bufio.NewScanner(file)
	// optionally, resize scanner's capacity for lines over 64K, see next example
	for scanner.Scan() {
		line := scanner.Text()
		fmt.Println(line)
		if len(line) == 0 {
			continue
		}
		i, err := strconv.Atoi(line)
		if err != nil {
			log.Fatal(err)
		}
		if prev >= 0 && i > prev {
			fmt.Println("--> increased")
			count++
		}
		prev = i
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return count
}

func puzzle2(fname string) int {
	file, err := os.Open(fname)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	var nn []int
	prev := 0
	count := 0
	scanner := bufio.NewScanner(file)
	// optionally, resize scanner's capacity for lines over 64K, see next example
	for scanner.Scan() {
		line := scanner.Text()
		fmt.Println(line)
		if len(line) == 0 {
			continue
		}
		n, err := strconv.Atoi(line)
		if err != nil {
			log.Fatal(err)
		}
		nn = append(nn, n)
		if len(nn) < 4 {
			prev += n
			continue
		}

		curr := prev - nn[len(nn)-4] + n
		fmt.Println("-->", prev, "-", nn[len(nn)-4], "+", n, " == ", curr)
		if curr > prev {
			fmt.Println("  --> increased")
			count++
		}
		prev = curr
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return count
}
