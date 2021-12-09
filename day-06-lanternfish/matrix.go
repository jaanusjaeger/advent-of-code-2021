package main

import (
	"fmt"
)

type matrix struct {
	data [][]int
}

func newMatrix(m, n int) *matrix {
	var data [][]int
	for i := 0; i < m; i++ {
		var row []int
		for j := 0; j < n; j++ {
			row = append(row, 0)
		}
		data = append(data, row)
	}
	return &matrix{data: data}
}

func (b *matrix) print() {
	for _, row := range b.data {
		for _, j := range row {
			fmt.Print(j, " ")
		}
		fmt.Println()
	}
}
