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

func (m *matrix) addRow(row []int) *matrix {
	m.data = append(m.data, row)
	return m
}

func (m *matrix) print() {
	for _, row := range m.data {
		for _, d := range row {
			fmt.Print(d, " ")
		}
		fmt.Println()
	}
}

func (m *matrix) printTr(stringer func(d int) string) {
	for _, row := range m.data {
		for _, d := range row {
			fmt.Print(stringer(d), " ")
		}
		fmt.Println()
	}
}
