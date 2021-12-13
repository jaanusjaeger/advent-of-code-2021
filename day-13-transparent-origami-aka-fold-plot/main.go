package main

import (
	"fmt"
	"strings"
)

func main() {
	file := openFileFromArgs()
	defer file.Close()

	plot := plot{}
	instructions := []instruction{}
	scanner := func(line string) error {
		ss := strings.Split(line, ",")
		if len(ss) == 2 {
			x := parseInt(ss[0])
			y := parseInt(ss[1])
			plot.dots = append(plot.dots, dot{x: x, y: y})
		} else {
			instr := strings.Split(line, "=")
			if len(instr) != 2 {
				return fmt.Errorf("invalid input: %s", line)
			}
			up := instr[0][len(instr[0])-1] == 'y'
			at := parseInt(instr[1])
			instructions = append(instructions, instruction{up: up, at: at})
		}
		return nil
	}
	scanLines(file, scanner)

	fmt.Println("INPUT:")
	fmt.Println(plot)
	fmt.Println(instructions)

	result1 := puzzle1(plot, instructions)
	result2 := puzzle2(plot, instructions)

	fmt.Println("RESULT1:", result1)
	fmt.Println("RESULT2:", result2)
}

func puzzle1(p plot, inst []instruction) int {
	result := 0

	p2 := p.fold(inst[0])
	result = len(p2.dots)

	return result
}

func puzzle2(p plot, inst []instruction) int {
	result := 0
	p2 := p
	for _, i := range inst {
		p2 = p2.fold(i)
	}
	p2.print()

	return result
}

type instruction struct {
	up bool
	at int
}

// plot is list of dots on a coordinate scale
type plot struct {
	dots []dot
}

func (p plot) fold(inst instruction) plot {
	if inst.up {
		return p.foldUp(inst.at)
	}
	return p.foldLeft(inst.at)
}

func (p plot) foldUp(y int) plot {
	fmt.Println("  foldUp", y)
	newDots := []dot{}
	keys := map[string]bool{}
	for _, d := range p.dots {
		newd := d.foldUp(y)
		if keys[newd.key()] {
			continue
		}
		keys[newd.key()] = true
		newDots = append(newDots, newd)
	}
	return plot{dots: newDots}
}

func (p plot) foldLeft(x int) plot {
	fmt.Println("  foldLeft", x)
	newDots := []dot{}
	keys := map[string]bool{}
	for _, d := range p.dots {
		newd := d.foldLeft(x)
		if keys[newd.key()] {
			continue
		}
		keys[newd.key()] = true
		newDots = append(newDots, newd)
	}
	return plot{dots: newDots}
}

func (p plot) print() {
	maxx := 0
	for _, d := range p.dots {
		maxx = max(maxx, d.x)
	}
	maxy := 0
	for _, d := range p.dots {
		maxy = max(maxy, d.y)
	}
	m := newMatrix(maxy+1, maxx+1)
	for _, d := range p.dots {
		m.data[d.y][d.x] = 1
	}
	m.printTr(func(d int) string {
		if d == 0 {
			return " "
		}
		return "#"
	})
}

// dot is a single dot with coordinates
type dot struct {
	x int
	y int
}

func (d dot) key() string {
	return fmt.Sprintf("%d,%d", d.x, d.y)
}

func (d dot) isRight(x int) bool {
	return d.x > x
}

func (d dot) isBelow(y int) bool {
	return d.y > y
}

func (d dot) foldUp(y int) dot {
	fmt.Print("    dot.foldUp ", y, d)
	if !d.isBelow(y) {
		fmt.Println(" --> no change")
		return dot{x: d.x, y: d.y}
	}
	newy := y - (d.y - y)
	fmt.Println(" --> ", d.x, newy)
	return dot{x: d.x, y: newy}
}

func (d dot) foldLeft(x int) dot {
	fmt.Print("    dot.foldLeft ", x, d)
	if !d.isRight(x) {
		fmt.Println(" --> no change")
		return dot{x: d.x, y: d.y}
	}
	newx := x - (d.x - x)
	fmt.Println(" --> ", newx, d.y)
	return dot{x: newx, y: d.y}
}
