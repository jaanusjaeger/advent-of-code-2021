package main

import (
	"fmt"
	"log"
	"strings"
)

func main() {
	file := openFileFromArgs()
	defer file.Close()

	i := 0
	var input []int
	var boards []*board
	scanner := func(line string) error {
		var row []int

		if i == 0 {
			splits := strings.Split(line, ",")
			for _, s := range splits {
				row = append(row, parseInt(strings.TrimSpace(s)))
			}

			input = row
		} else {
			splits := strings.Split(line, " ")
			for _, s := range splits {
				if s != "" {
					row = append(row, parseInt(strings.TrimSpace(s)))
				}
			}

			var b *board
			if len(boards) == 0 || boards[len(boards)-1].isCreated() {
				b = &board{}
				boards = append(boards, b)
			} else {
				b = boards[len(boards)-1]
			}
			b.addRow(row)
		}

		i++

		return nil
	}
	scanLines(file, scanner)

	fmt.Println()
	fmt.Println("INPUT:", input)
	fmt.Println("BOARDS:")
	printBoards(boards)

	fmt.Println()
	fmt.Println("PLAY")

	result1 := puzzle1(input, boards)
	result2 := puzzle2(input, boards)

	fmt.Println("RESULT1:", result1)
	fmt.Println("RESULT2:", result2)
}

func puzzle1(input []int, boards []*board) int {
	left := 0
	final := 0
outer:
	for _, n := range input {
		for _, b := range boards {
			b.mark(n)
			if b.hasWon() {
				left = b.left
				final = n
				break outer
			}
		}
	}

	return left * final
}

func puzzle2(input []int, boards []*board) int {
	left := 0
	final := 0
outer:
	for _, n := range input {
		fmt.Println("-- PICKED:", n)
		correction := 0
		for ind, b := range boards {
			i := ind - correction
			b.mark(n)
			if b.hasWon() {
				fmt.Println("---- BOARD WON:", i)
				if len(boards) == 1 {
					left = b.left
					final = n
					break outer
				}
				var boards2 []*board
				for j, b := range boards {
					if j != i {
						boards2 = append(boards2, b)
					}
				}
				boards = boards2
				correction++
				fmt.Println("---- LEFT BOARDS:", len(boards))
				//printBoards(boards)
			}
		}
	}

	return left * final
}

func printBoards(boards []*board) {
	for _, b := range boards {
		for _, row := range b.board {
			for _, j := range row {
				fmt.Print(j, " ")
			}
			fmt.Println()
		}
		fmt.Println("LEFT:", b.left)
		fmt.Println()
	}
}

type board struct {
	board [][]int
	left  int
}

func (b *board) isCreated() bool {
	return len(b.board) == 5
}

func (b *board) addRow(row []int) {
	if b.isCreated() {
		log.Fatal("Board already created", b)
	}
	if len(row) != 5 {
		log.Fatal("Invalid row - must have length 5", row)
	}
	b.board = append(b.board, row)
	for _, n := range row {
		b.left += n
	}
}

func (b *board) mark(k int) {
	for i, row := range b.board {
		for j, n := range row {
			if n == k {
				b.board[i][j] = -1
				b.left -= k
				return
			}
		}
	}
}

func (b *board) hasWon() bool {
	for i := range b.board {
		if b.checkRowWon(i) || b.checkColWon(i) {
			return true
		}
	}

	return false
}

func (b *board) checkRowWon(i int) bool {
	for j := range b.board {
		if b.board[i][j] != -1 {
			return false
		}
	}
	return true
}

func (b *board) checkColWon(j int) bool {
	for i := range b.board {
		if b.board[i][j] != -1 {
			return false
		}
	}
	return true
}
