package main

import (
	"fmt"
	"log"
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

	size := 1000
	board := newBoard(size, size)
	scanner := func(line string) error {
		points := strings.Split(line, " -> ")
		if len(points) != 2 {
			log.Fatal("Invalid input line:", line)
		}
		start := strings.Split(points[0], ",")
		if len(start) != 2 {
			log.Fatal("Invalid start coord:", points[0])
		}
		end := strings.Split(points[1], ",")
		if len(end) != 2 {
			log.Fatal("Invalid end coord:", points[1])
		}
		x1 := parseInt(start[0])
		y1 := parseInt(start[1])
		x2 := parseInt(end[0])
		y2 := parseInt(end[1])

		board.mark1(x1, y1, x2, y2)
		return nil
	}
	scanLines(file, scanner)

	fmt.Println()
	fmt.Println("BOARD:")
	printBoard(board)

	return board.countAbove(1)
}

func puzzle2() int {
	file := openFileFromArgs()
	defer file.Close()

	size := 1000
	board := newBoard(size, size)
	scanner := func(line string) error {
		points := strings.Split(line, " -> ")
		if len(points) != 2 {
			log.Fatal("Invalid input line:", line)
		}
		start := strings.Split(points[0], ",")
		if len(start) != 2 {
			log.Fatal("Invalid start coord:", points[0])
		}
		end := strings.Split(points[1], ",")
		if len(end) != 2 {
			log.Fatal("Invalid end coord:", points[1])
		}
		x1 := parseInt(start[0])
		y1 := parseInt(start[1])
		x2 := parseInt(end[0])
		y2 := parseInt(end[1])

		board.mark2(x1, y1, x2, y2)
		return nil
	}
	scanLines(file, scanner)

	fmt.Println()
	fmt.Println("BOARD:")
	printBoard(board)

	return board.countAbove(1)
}

func printBoard(b *board) {
	for _, row := range b.board {
		for _, j := range row {
			fmt.Print(j, " ")
		}
		fmt.Println()
	}
}

type board struct {
	board [][]int
}

func newBoard(m, n int) *board {
	var data [][]int
	for i := 0; i < m; i++ {
		var row []int
		for j := 0; j < n; j++ {
			row = append(row, 0)
		}
		data = append(data, row)
	}
	return &board{board: data}
}

func (b *board) mark1(x1, y1, x2, y2 int) {
	if x1 != x2 && y1 != y2 {
		// Diagonal
		return
	}
	// xs and ys are the start coords and xe and ye are the end coords from upper left to lower right
	var xs, ys, xe, ye int
	if x1 == x2 {
		xs = x1
		xe = x1
		ys = min(y1, y2)
		ye = max(y1, y2)
	} else {
		xs = min(x1, x2)
		xe = max(x1, x2)
		ys = y1
		ye = y1
	}
	for i := xs; i <= xe; i++ {
		for j := ys; j <= ye; j++ {
			b.board[i][j] += 1
		}
	}
}

func (b *board) mark2(x1, y1, x2, y2 int) {
	if x1 != x2 && y1 != y2 && (x1+y1 != x2+y2) && (max(x1, x2)-min(x1, x2) != max(y1, y2)-min(y1, y2)) {
		fmt.Printf("not considered: %d, %d -> %d, %d\n", x1, y1, x2, y2)
		return
	}

	length := max(max(x1, x2)-min(x1, x2), max(y1, y2)-min(y1, y2))
	fmt.Printf("Length: %d, %d -> %d, %d == %d\n", x1, y1, x2, y2, length)

	stepx := 0
	if x1 < x2 {
		stepx = 1
	} else if x1 > x2 {
		stepx = -1
	}
	stepy := 0
	if y1 < y2 {
		stepy = 1
	} else if y1 > y2 {
		stepy = -1
	}
	x := x1
	y := y1
	for i := 0; i <= length; i++ {
		b.board[x][y] += 1
		x += stepx
		y += stepy
	}
}

func (b *board) countAbove(k int) int {
	count := 0
	for _, row := range b.board {
		for _, cell := range row {
			if cell > k {
				count++
			}
		}
	}
	return count
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
