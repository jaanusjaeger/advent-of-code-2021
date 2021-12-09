package main

import (
	"fmt"
	"strings"
)

func main() {
	file := openFileFromArgs()
	defer file.Close()

	crabs := []int{}
	scanner := func(line string) error {
		ss := strings.Split(line, ",")
		for _, s := range ss {
			n := parseInt(s)
			crabs = append(crabs, n)
		}
		return nil
	}
	scanLines(file, scanner)

	fmt.Println("INPUT:")
	fmt.Println(crabs)

	result1 := puzzle1(crabs)
	result2 := puzzle2(crabs)

	fmt.Println("RESULT1:", result1)
	fmt.Println("RESULT2:", result2)
}

func puzzle1(crabs []int) int {
	min := -1

	for _, p := range crabs {
		diff := 0
		for _, c := range crabs {
			diff += abs(p - c)
		}
		if min < 0 || diff < min {
			fmt.Println("POS:", p, "FUEL:", diff)
			min = diff
		}
	}

	result := min

	return result
}

func puzzle2(crabs []int) int {
	maxCrab := -1
	for _, c := range crabs {
		if maxCrab < 0 || c > maxCrab {
			maxCrab = c
		}
	}
	fmt.Println("MAX CRAB:", maxCrab)

	min := -1

	for p := 0; p <= maxCrab; p++ {
		diff := 0
		//fmt.Println()
		for _, c := range crabs {
			f := fuel(p, c)
			diff += f
			//fmt.Println(c, "->", p, "=", f, " -- total:", diff)
		}
		if min < 0 || diff < min {
			fmt.Println("POS:", p, "FUEL:", diff)
			min = diff
		}
	}

	result := min

	return result
}

func fuel(x, y int) int {
	a := abs(x - y)
	/*
		return ((1.0 + a) / 2.0) * a
	*/
	s := 0
	for i := 1; i <= a; i++ {
		s += i
	}
	return s
}
