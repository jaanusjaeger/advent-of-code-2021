package main

import (
	"fmt"
	"strings"
)

func main() {
	file := openFileFromArgs()
	defer file.Close()

	lines := []string{}
	nodes := nodes{}
	scanner := func(line string) error {
		lines = append(lines, line)
		nodes = append(nodes, parse(line))
		return nil
	}
	scanLines(file, scanner)

	fmt.Println("INPUT:")
	fmt.Println(nodes)

	// TODO debugging
	// for _, n := range nodes {
	// 	if reduce(n) {
	// 		fmt.Println("INPUT NOT REDUCED")
	// 	}
	// }

	result1 := puzzle1(nodes)
	result2 := puzzle2(lines)

	fmt.Println("RESULT1:", result1)
	fmt.Println("RESULT2:", result2)
}

func puzzle1(nn nodes) string {
	root := nn[0]
	for _, n := range nn[1:] {
		root = &pair{left: root, right: n}
		reduce(root)
	}

	result := fmt.Sprint(root.Magnitude())

	return result
}

func puzzle2(lines []string) string {
	maxMag := 0

	for _, l1 := range lines {
		for _, l2 := range lines {
			if l1 == l2 {
				continue
			}
			n1 := parse(l1)
			n2 := parse(l2)
			sum := &pair{left: n1, right: n2}
			reduce(sum)
			mag := sum.Magnitude()
			maxMag = max(maxMag, mag)
		}
	}

	result := fmt.Sprint(maxMag)

	return result
}

func parse(s string) node {
	pairs := stack{}

	setChild := func(n node) {
		if pairs.size == 0 {
			pairs.push(n.(*pair))
			return
		}
		if p, ok := n.(*pair); ok {
			p.parent = n
		}
		if pairs.peek().left == nil {
			pairs.peek().left = n
		} else {
			pairs.peek().right = n
		}
	}

	for i := 0; i < len(s); i++ {
		c := s[i]
		// fmt.Println("  read", string(c), "number of pairs:", pairs.size)
		if c == '[' {
			// fmt.Println("    --> start new pair")
			p := pair{}
			setChild(&p)
			pairs.push(&p)
		} else if c == ']' {
			// fmt.Println("    --> end pair")
			pairs.pop()
		} else if c == ',' {
			// fmt.Println("    --> comma, nop")
			continue
		} else {
			// fmt.Println("    --> number")
			j := strings.IndexFunc(s[i:], func(r rune) bool { return r == ',' || r == ']' })
			n := number{n: parseInt(s[i : i+j])}
			i += j - 1
			setChild(&n)
		}
	}

	return pairs.pop()
}

func reduce(n node) (result bool) {
	fmt.Println("REDUCING:", n)
	for {
		last := []*number{nil}
		add := 0
		fmt.Println("    reduce: -->", n)
		if exploded, _ := n.Explode(0, last, &add); exploded {
			result = true
			continue
		}
		if split, _ := n.Split(); split {
			result = true
			continue
		}
		break
	}
	fmt.Println("REDUCED:", n)
	return result
}

type nodes []node

func (nn nodes) String() string {
	result := []string{}
	for _, n := range nn {
		result = append(result, fmt.Sprint(n))
	}
	return fmt.Sprintf("[\n%s\n]", strings.Join(result, ",\n"))
}

type node interface {
	Explode(depth int, last []*number, add *int) (bool, bool)
	Split() (bool, node)
	Magnitude() int
}

type number struct {
	n int
}

func (n *number) Explode(depth int, last []*number, add *int) (bool, bool) {
	// fmt.Println("        Explode (Number):", n, depth, last, *add)
	last[0] = n
	if *add != 0 {
		fmt.Println("          --> Explode (Number): ADD:", *add)
		n.n += *add
		*add = 0
		return true, true
	}
	return false, false
}

func (n *number) Split() (bool, node) {
	// fmt.Println("        Split (Number):", n)
	if n.n < 10 {
		return false, nil
	}
	fmt.Println("          --> Split (Number):", n)
	p := pair{
		left:  &number{n.n / 2},
		right: &number{n.n/2 + n.n%2},
	}
	return true, &p
}

func (n *number) Magnitude() int {
	return n.n
}

func (n number) String() string {
	return fmt.Sprintf("%d", n.n)
}

type pair struct {
	parent node
	left   node
	right  node
}

func (p *pair) Explode(depth int, last []*number, add *int) (bool, bool) {
	// fmt.Println("        Explode (Pair  ):", p, depth, last, *add)
	if depth > 3 && *add == 0 {
		fmt.Println("          --> Explode (Pair  ): YES:", p)
		lv := (p.left).(*number)
		rv := (p.right).(*number)
		if last[0] != nil {
			last[0].n += lv.n
		}
		*add = rv.n
		return true, false
	}
	leftExploded, leftDone := p.left.Explode(depth+1, last, add)
	if leftExploded && !leftDone {
		p.left = &number{}
	}
	if leftExploded && *add == 0 {
		return true, true
	}
	rightExploded, rightDone := p.right.Explode(depth+1, last, add)
	if rightExploded && !rightDone {
		p.right = &number{}
	}
	if rightExploded && *add == 0 {
		return true, true
	}
	return leftExploded || rightExploded, leftExploded || rightExploded
}

func (p *pair) Split() (bool, node) {
	// fmt.Println("        Split (Pair  ):", p)
	ls, lp := p.left.Split()
	if ls {
		if lp != nil {
			p.left = lp
		}
		return true, nil
	}
	rs, rp := p.right.Split()
	if rs {
		if rp != nil {
			p.right = rp
		}
		return true, nil
	}
	return false, nil
}

func (p *pair) Magnitude() int {
	return 3*p.left.Magnitude() + 2*p.right.Magnitude()
}

func (p pair) String() string {
	return fmt.Sprintf("[%v,%v]", p.left, p.right)
}
