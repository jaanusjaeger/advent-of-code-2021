package main

import (
	"fmt"
	"sort"
	"strings"
)

var allNodes map[string]*matrixNode

func main() {
	file := openFileFromArgs()
	defer file.Close()

	m := &matrix{}
	scanner := func(line string) error {
		var row []int

		ss := strings.Split(line, "")
		for _, s := range ss {
			row = append(row, parseInt(strings.TrimSpace(s)))
		}

		m.addRow(row)
		return nil
	}
	scanLines(file, scanner)

	fmt.Println("INPUT:")
	m.print()

	result1 := puzzle1(m)
	result2 := puzzle2(m)

	fmt.Println("RESULT1:", result1)
	fmt.Println("RESULT2:", result2)
}

func puzzle1(m *matrix) string {
	allNodes = make(map[string]*matrixNode)
	for i, row := range m.data {
		for j := range row {
			allNodes[matrixNodeID(i, j)] = &matrixNode{m: m, i: i, j: j}
		}
	}

	size := len(m.data)
	pq := &sortedSlicePriorityQueue{}
	start := allNodes[matrixNodeID(0, 0)]
	destID := matrixNodeID(size-1, size-1)

	// This is needed for 'dijkstra' function to not enter loops
	start.prev = start
	fmt.Println("START:", start)
	for _, n := range start.Neighbors() {
		fmt.Println("  neighbor:", n)
	}

	pq.Push(start)
	fmt.Println("PQ:", pq.elem)
	dest := dijkstra(pq, destID)

	fmt.Println("PATH (reversed):")
	for d := dest; d != nil && d != d.Prev(); d = d.Prev() {
		fmt.Println("  ", d.ID(), "-", d.Dist())
	}

	return fmt.Sprint(dest.Dist())
}

func puzzle2(m *matrix) string {
	size := len(m.data)
	size5 := 5 * size

	mm := newMatrix(size5, size5)
	for i := 0; i < size5; i++ {
		rowTile := i / size
		for j := 0; j < size5; j++ {
			colTile := j / size
			add := rowTile + colTile
			d := m.data[i%size][j%size]
			val := d + add
			if val > 9 {
				val = val % 9
			}
			mm.data[i][j] = val
		}
	}

	// m.print()
	// mm.print()

	return puzzle1(mm)
}

func dijkstra(pq PriorityQueue, destID string) Node {
	for {
		node := pq.Pop()
		fmt.Println("  REACHED:", node)
		if node.ID() == destID {
			return node
		}
		for _, n := range node.Neighbors() {
			// NB! In normal graph a 'D-tour' may be shortcut, but not in this graph!!!
			if n.Prev() == nil {
				n.SetPrev(node)
				pq.Push(n)
			}
		}
	}
}

type Node interface {
	ID() string
	Neighbors() []Node
	SetPrev(Node)
	Dist() int
	Prev() Node
}

type matrixNode struct {
	m    *matrix
	i    int
	j    int
	prev Node
	dist int
}

func (n *matrixNode) ID() string {
	return matrixNodeID(n.i, n.j)
}

func (n *matrixNode) Neighbors() []Node {
	result := []Node{}

	if n.i > 0 {
		result = append(result, allNodes[matrixNodeID(n.i-1, n.j)])
	}
	if n.i < len(n.m.data)-1 {
		result = append(result, allNodes[matrixNodeID(n.i+1, n.j)])
	}
	if n.j > 0 {
		result = append(result, allNodes[matrixNodeID(n.i, n.j-1)])
	}
	if n.j < len(n.m.data)-1 {
		result = append(result, allNodes[matrixNodeID(n.i, n.j+1)])
	}

	return result
}

func (n *matrixNode) SetPrev(p Node) {
	n.prev = p
	n.dist = p.Dist() + n.m.data[n.i][n.j]
}

func (n *matrixNode) Dist() int {
	return n.dist
}

func (n *matrixNode) Prev() Node {
	return n.prev
}

func (n *matrixNode) String() string {
	return fmt.Sprintf("%s (%d)", n.ID(), n.Dist())
}

func matrixNodeID(i, j int) string {
	return fmt.Sprintf("[%d,%d]", i, j)
}

type PriorityQueue interface {
	Push(Node)
	Pop() Node
}

type sortedSlicePriorityQueue struct {
	elem []Node
}

func (pq *sortedSlicePriorityQueue) Push(e Node) {
	//defer func() { fmt.Println("    pq.Push() ->", pq.elem) }()
	i := sort.Search(len(pq.elem), func(i int) bool {
		return pq.elem[i].Dist() >= e.Dist()
	})
	if i == len(pq.elem) {
		pq.elem = append(pq.elem, e)
		return
	}
	pq.elem = append(pq.elem[:i+1], pq.elem[i:]...)
	pq.elem[i] = e
}

func (pq *sortedSlicePriorityQueue) Pop() Node {
	//defer func() { fmt.Println("    pq.Pop() ->", pq.elem) }()
	if len(pq.elem) == 0 {
		return nil
	}
	e := pq.elem[0]

	pq.elem = pq.elem[1:]

	return e
}
