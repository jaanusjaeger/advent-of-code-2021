package main

import (
	"fmt"
	"strings"
)

func main() {
	file := openFileFromArgs()
	defer file.Close()

	o := &octopuses{}
	scanner := func(line string) error {
		var row []int

		ss := strings.Split(line, "")
		for _, s := range ss {
			if s != "" {
				row = append(row, parseInt(strings.TrimSpace(s)))
			}
		}

		o.addRow(row)
		return nil
	}
	scanLines(file, scanner)

	fmt.Println("INPUT:")
	o.print()

	// Assume that puzzle2 day comes after 100
	puzzle1Days := 100
	result1 := puzzle1(o, puzzle1Days)
	result2 := puzzle2(o) + puzzle1Days

	fmt.Println("RESULT1:", result1)
	fmt.Println("RESULT2:", result2)
}

func puzzle1(o *octopuses, days int) int {
	count := 0

	for d := 0; d < days; d++ {
		fmt.Println("DAY ", d+1, ":")
		o.initFlashed()
		// 1. Increase energy
		for i, row := range o.data {
			for j, e := range row {
				o.data[i][j] = e + 1
			}
		}
		// 2. Flash
		for i, row := range o.data {
			for j := range row {
				count += o.flash(i, j, false)
			}
		}
		// 3. Set flashed to 0
		for i, row := range o.data {
			for j, e := range row {
				if e >= 10 {
					o.data[i][j] = 0
				}
			}
		}

		o.print()
	}

	result := count
	return result
}

func puzzle2(o *octopuses) int {
	day := 1
	for ; ; day++ {
		fmt.Println("DAY ", day+1, ":")
		o.initFlashed()
		count := 0
		// 1. Increase energy
		for i, row := range o.data {
			for j, e := range row {
				o.data[i][j] = e + 1
			}
		}
		// 2. Flash
		for i, row := range o.data {
			for j := range row {
				count += o.flash(i, j, false)
			}
		}
		// 3. Set flashed to 0
		for i, row := range o.data {
			for j, e := range row {
				if e >= 10 {
					o.data[i][j] = 0
				}
			}
		}

		o.print()

		if count == len(o.data)*len(o.data[0]) {
			break
		}
	}

	result := day
	return result
}

type octopuses struct {
	matrix
	flashed [][]bool
}

func (o *octopuses) initFlashed() {
	o.flashed = [][]bool{}
	for _, row := range o.data {
		o.flashed = append(o.flashed, make([]bool, len(row)))
	}
}

func (o *octopuses) flash(i, j int, increase bool) int {
	// Detect out-of-bounds
	if i < 0 || i >= len(o.data) || j < 0 || j >= len(o.data[0]) {
		return 0
	}
	// Consume the flash
	if increase {
		o.data[i][j] += 1
	}
	// Either not subject to flashing or already flashed
	if o.data[i][j] < 10 || o.flashed[i][j] {
		return 0
	}
	o.flashed[i][j] = true
	fmt.Println("  flash at", i, j)
	return 1 +
		o.flash(i-1, j-1, true) +
		o.flash(i-1, j, true) +
		o.flash(i-1, j+1, true) +
		o.flash(i, j-1, true) +
		o.flash(i, j+1, true) +
		o.flash(i+1, j-1, true) +
		o.flash(i+1, j, true) +
		o.flash(i+1, j+1, true)
}

func (o *octopuses) countFlashed() int {
	count := 0
	for _, row := range o.flashed {
		for _, v := range row {
			if v {
				count++
			}
		}
	}
	return count
}
