package main

import (
	"fmt"
	"strings"
)

func main() {
	result1 := puzzle1()
	result2 := puzzle2()

	fmt.Println("RESULT1:", result1)
	fmt.Println("RESULT2:", result2)
}

func puzzle1() int {
	file := openFileFromArgs()
	defer file.Close()

	fish := []int{}
	days := 80
	birthValue := 8
	resetValue := 6
	scanner := func(line string) error {
		ss := strings.Split(line, ",")
		for _, s := range ss {
			fish = append(fish, parseInt(s))
		}
		return nil
	}
	scanLines(file, scanner)

	fmt.Println("INITIAL:")
	fmt.Println(fish)

	for d := 0; d < days; d++ {
		for i, f := range fish {
			if f == 0 {
				fish[i] = resetValue
				fish = append(fish, birthValue)
			} else {
				fish[i] -= 1
			}
		}
		fmt.Println("END OF DAY", d)
		//fmt.Println(fish)
	}

	result := len(fish)

	return result
}

func puzzle2() int {
	file := openFileFromArgs()
	defer file.Close()

	//fish := []int{}
	days := 256
	resetValue := 6
	birthValue := 8
	ages := [9]int{}
	scanner := func(line string) error {
		ss := strings.Split(line, ",")
		for _, s := range ss {
			n := parseInt(s)
			//fish = append(fish, n)
			/*
				if a, ok := ages[n]; ok {
					ages
				}
			*/
			ages[n] += 1
		}
		return nil
	}
	scanLines(file, scanner)

	fmt.Println("INITIAL:")
	fmt.Println(ages)

	for d := 0; d < days; d++ {
		var births int
		for a, n := range ages {
			if a == 0 {
				births = n
				ages[a] = 0
			} else {
				ages[a] = 0
				ages[a-1] += n
			}
		}
		ages[resetValue] += births
		ages[birthValue] += births
		fmt.Println("END OF DAY", d)
		fmt.Println(ages)
	}

	//result := len(fish)
	result := 0
	for _, a := range ages {
		result += a
	}

	return result
}
