package main

import (
	"fmt"
	"strings"
)

func main() {
	file := openFileFromArgs()
	defer file.Close()

	// NB! Making simplification due to input data:
	// *) enhancement algorithm turns 3x3 of '.' into '#' and 3x3 of '#' into '.'
	// *) after even number of steps all "infinite" edges of '#' will be '.' again
	// *) Add sufficient (finite) buffer around the initial matrix, to preserve all
	//    non-infinite enhancements.

	// Add 'extention' to the beginning and end of each row, also add 'bufferSize' empty rows before and after
	bufferSize := 60
	extention := make([]int, bufferSize)
	enhancementAlg := ""
	m := &matrix{}
	scanner := func(line string) error {
		if enhancementAlg == "" {
			enhancementAlg = line
			return nil
		}

		var row []int
		row = append(row, extention...)

		ss := strings.Split(line, "")
		// Add rows before
		if len(m.data) == 0 {
			m.addNTimesEmptyRow(bufferSize, len(ss)+2*bufferSize)
		}
		for _, s := range ss {
			v := 0
			if s == "#" {
				v = 1
			}
			row = append(row, v)
		}

		row = append(row, extention...)

		m.addRow(row)
		return nil
	}
	scanLines(file, scanner)

	// Add empty rows to the beginning and end
	m.addNTimesEmptyRow(bufferSize, len(m.data[0]))

	fmt.Println("INPUT:")
	fmt.Println(enhancementAlg)
	m.print()

	result1 := puzzle1(enhancementAlg, m, 2)
	result2 := puzzle1(enhancementAlg, m, 50)

	fmt.Println("RESULT1:", result1)
	fmt.Println("RESULT2:", result2)
}

func puzzle1(enhancementAlg string, m *matrix, iterations int) string {
	for it := 0; it < iterations; it++ {
		fmt.Println("Iteration:", it)
		m = enhance(m, enhancementAlg)
		// m.print()
	}
	m.print()

	count := 0
	for _, row := range m.data {
		for _, d := range row {
			if d > 0 {
				count++
			}
		}
	}

	result := fmt.Sprint(count)

	return result
}

func puzzle2() string {
	result := ""

	return result
}

func enhance(m *matrix, alg string) *matrix {
	l := len(m.data)
	cpy := m.cpy()

	for i := 1; i < l-1; i++ {
		row := m.data[i]
		for j := 1; j < len(row)-1; j++ {
			pixelCode := m.getPixelCode(i, j)
			v := 0
			if alg[pixelCode] == '#' {
				v = 1
			}
			cpy.data[i][j] = v
		}
	}

	// Since the edge bits are ignored in the calculation, set the edges according to any bit
	// that is calculated as infinite bit - bit right next to the edge
	infiniteBit := cpy.data[1][1]
	for i := 0; i < l; i++ {
		cpy.data[0][i] = infiniteBit
		cpy.data[l-1][i] = infiniteBit
		cpy.data[i][0] = infiniteBit
		cpy.data[i][l-1] = infiniteBit
	}

	return cpy
}

func (m *matrix) addNTimesEmptyRow(n int, l int) {
	for i := 0; i < n; i++ {
		row := make([]int, l)
		m.addRow(row)
	}
}

func (m *matrix) cpy() *matrix {
	c := matrix{}

	for _, row := range m.data {
		rowCpy := make([]int, len(row))
		copy(rowCpy, row)
		c.addRow(rowCpy)
	}

	return &c
}

func (m *matrix) getPixelCode(i, j int) int {
	v := 0
	for ii := -1; ii < 2; ii++ {
		for jj := -1; jj < 2; jj++ {
			// fmt.Println("    getPixelCode", i, j, " >>", m.data[i+ii][j+jj])
			v *= 2
			v += m.data[i+ii][j+jj]
		}
	}
	return v
}
