package main

import (
	"fmt"
	"sort"
	"strings"
)

func main() {
	file := openFileFromArgs()
	defer file.Close()

	caves := map[string]*cave{}
	scanner := func(line string) error {
		ss := strings.Split(line, "-")
		if len(ss) != 2 {
			return fmt.Errorf("invalid input: %s", line)
		}
		sname := ss[0]
		start, ok := caves[sname]
		if !ok {
			start = &cave{name: sname}
			caves[sname] = start
		}
		ename := ss[1]
		end, ok := caves[ename]
		if !ok {
			end = &cave{name: ename}
			caves[ename] = end
		}
		start.neighbours = append(start.neighbours, end)
		end.neighbours = append(end.neighbours, start)
		return nil
	}
	scanLines(file, scanner)

	// Sort to match example solution and simplify debugging
	for _, c := range caves {
		sorter := func(i, j int) bool {
			return c.neighbours[i].name < c.neighbours[j].name
		}
		sort.Slice(c.neighbours, sorter)
	}

	fmt.Println("INPUT:")
	fmt.Println(caves)

	result1 := puzzle1(caves)
	result2 := puzzle2(caves)

	fmt.Println("RESULT1:", result1)
	fmt.Println("RESULT2:", result2)
}

func puzzle1(caves map[string]*cave) int {
	result := 0

	start := caves["start"]
	guard := func(cave *cave, p path) bool {
		return !cave.isSmall() || !p.contains(cave)
	}
	result = scan(start, path{}, guard)

	return result
}

func puzzle2(caves map[string]*cave) int {
	result := 0

	fmt.Println()
	start := caves["start"]
	guard := func(cave *cave, p path) bool {
		if !cave.isSmall() {
			return true
		}
		if cave.name == "start" && len(p.caves) > 0 {
			return false
		}
		smalls := map[string]int{cave.name: 1}
		twiceIsUsed := false
		for _, c := range p.caves {
			if c.isSmall() && c.name != "start" {
				if smalls[c.name] > 0 {
					if twiceIsUsed {
						return false
					}
					twiceIsUsed = true
				}
				smalls[c.name]++
			}
		}
		return true
	}
	result = scan(start, path{}, guard)

	return result
}

func scan(cave *cave, p path, guard func(*cave, path) bool) int {
	fmt.Println("  visit", cave, "small:", cave.isSmall(), "path:", p)
	if cave.name == "end" {
		fmt.Println("    END")
		return 1
	}
	if !guard(cave, p) {
		return 0
	}

	p.add(cave)
	count := 0
	for _, c := range cave.neighbours {
		count += scan(c, p, guard)
	}

	return count
}

type cave struct {
	name       string
	neighbours []*cave
}

func (c *cave) String() string {
	names := []string{}
	for _, n := range c.neighbours {
		names = append(names, n.name)
	}
	return fmt.Sprintf("<%s -> %s>", c.name, names)
}

func (c *cave) isSmall() bool {
	return c.name[0] >= 'a' && c.name[0] <= 'z'
}

type path struct {
	caves []*cave
}

func (p *path) add(cave *cave) {
	p.caves = append(p.caves, cave)
}

func (p *path) contains(cave *cave) bool {
	for _, c := range p.caves {
		if c == cave {
			return true
		}
	}
	return false
}

func (p path) String() string {
	names := []string{}
	for _, n := range p.caves {
		names = append(names, n.name)
	}
	return fmt.Sprintf("%s", names)
}
