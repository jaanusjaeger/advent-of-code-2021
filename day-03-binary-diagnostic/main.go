package main

import (
	"fmt"
	"log"
)

func main() {
	file := openFileFromArgs()
	defer file.Close()

	nn := []string{}
	scanner := func(line string) error {
		intFromBinaryString(line)
		nn = append(nn, line)
		return nil
	}
	scanLines(file, scanner)

	result1 := puzzle1(nn)
	result2 := puzzle2(nn)

	fmt.Println("RESULT1:", result1)
	fmt.Println("RESULT2:", result2)
}

func puzzle1(nn []string) int {
	common := ""

	for i := range nn[0] {
		count1 := 0
		count0 := 0
		for _, n := range nn {
			if n[i] == '1' {
				count1++
			} else {
				count0++
			}
		}
		if count1 > count0 {
			common += "1"
		} else if count1 < count0 {
			common += "0"
		} else {
			log.Fatal("eve popularity", i)
		}
	}

	var gammaStr string
	var epsilonStr string

	gammaStr = common
	for _, c := range gammaStr {
		if c == '1' {
			epsilonStr += "0"
		} else {
			epsilonStr += "1"
		}
	}

	gamma := intFromBinaryString(gammaStr)
	epsilon := intFromBinaryString(epsilonStr)

	return gamma * epsilon
}

func puzzle2(nn []string) int {
	oxygenStr := filter(nn, 0, false)
	scrubberStr := filter(nn, 0, true)
	oxygen := intFromBinaryString(oxygenStr)
	scrubber := intFromBinaryString(scrubberStr)

	return oxygen * scrubber
}

func filter(nn []string, p int, flip bool) string {
	var sample byte
	count1 := 0
	for _, n := range nn {
		if n[p] == '1' {
			count1++
		}
	}
	count0 := len(nn) - count1
	if count1 >= count0 {
		if !flip {
			sample = '1'
		} else {
			sample = '0'
		}
	} else if count1 < count0 {
		if !flip {
			sample = '0'
		} else {
			sample = '1'
		}
	}

	fmt.Println("----> TARGET", sample)

	candidates := []string{}
	for _, n := range nn {
		if n[p] == sample {
			candidates = append(candidates, n)
		}
	}
	fmt.Println("CANDIDATES", sample, p, candidates)
	if len(candidates) == 1 {
		return candidates[0]
	}
	return filter(candidates, p+1, flip)
}
