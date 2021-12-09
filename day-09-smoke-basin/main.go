package main

import (
	"fmt"
	"sort"
	"strings"
)

func main() {
	file := openFileFromArgs()
	defer file.Close()

	c := &cave{}
	scanner := func(line string) error {
		var row []int

		ss := strings.Split(line, "")
		for _, s := range ss {
			if s != "" {
				row = append(row, parseInt(strings.TrimSpace(s)))
			}
		}

		c.addRow(row)
		return nil
	}
	scanLines(file, scanner)

	fmt.Println("INPUT:")
	c.print()

	result1 := puzzle1(c)
	result2 := puzzle2(c)

	fmt.Println("RESULT1:", result1)
	fmt.Println("RESULT2:", result2)
}

func puzzle1(c *cave) int {
	count := 0
	sum := 0
	for i, row := range c.data {
		for j, d := range row {
			if (i <= 0 || i > 0 && c.data[i-1][j] > d) &&
				(i >= len(c.data)-1 || i < len(c.data)-1 && c.data[i+1][j] > d) &&
				(j <= 0 || j > 0 && c.data[i][j-1] > d) &&
				(j >= len(row)-1 || j < len(row)-1 && c.data[i][j+1] > d) {
				fmt.Println("--> low point at:", i, j, "  count:", count)
				count++
				sum += d + 1
			}
		}
	}

	result := sum
	return result
}

func puzzle2(c *cave) int {
	c.init()

	basins := []int{}
	for i, row := range c.data {
		for j := range row {
			b := c.spread(i, j)
			if b > 0 {
				basins = append(basins, b)
				fmt.Println("--> basin at", i, j, " = ", b)
			}
		}
	}
	sort.Ints(basins)
	l := len(basins)

	result := basins[l-1] * basins[l-2] * basins[l-3]
	return result
}

type cave struct {
	matrix
	visited [][]bool
}

func (c *cave) init() {
	for _, row := range c.data {
		c.visited = append(c.visited, make([]bool, len(row)))
	}
}

func (c *cave) spread(i, j int) int {
	if i < 0 || i >= len(c.data) || j < 0 || j >= len(c.data[0]) {
		return 0
	}
	if c.visited[i][j] || c.data[i][j] == 9 {
		return 0
	}
	c.visited[i][j] = true
	return 1 +
		c.spread(i-1, j) +
		c.spread(i+1, j) +
		c.spread(i, j-1) +
		c.spread(i, j+1)
}
