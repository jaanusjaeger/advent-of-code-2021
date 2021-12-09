package main

import (
	"fmt"
	"sort"
	"strings"
)

func main() {
	file := openFileFromArgs()
	defer file.Close()

	entries := []entry{}
	scanner := func(line string) error {
		ss := strings.Split(line, "|")
		if len(ss) != 2 {
			return fmt.Errorf("invalid input: %s", line)
		}
		ps := strings.Split(strings.TrimSpace(ss[0]), " ")
		if len(ps) != 10 {
			return fmt.Errorf("invalid pattern: %s", ss[0])
		}
		os := strings.Split(strings.TrimSpace(ss[1]), " ")
		if len(os) != 4 {
			return fmt.Errorf("invalid output: %s", ss[1])
		}
		entries = append(entries, entry{ps, os})
		return nil
	}
	scanLines(file, scanner)

	fmt.Println("INPUT:")
	fmt.Println(entries)

	result1 := puzzle1(entries)
	result2 := puzzle2(entries)

	fmt.Println("RESULT1:", result1)
	fmt.Println("RESULT2:", result2)
}

func puzzle1(entries []entry) int {
	count := 0
	for _, e := range entries {
		for _, o := range e.output {
			switch len(o) {
			case 2, 3, 4, 7:
				count++
			}
		}
	}

	result := count
	return result
}

func puzzle2(entries []entry) int {
	sum := 0
	for _, e := range entries {
		sum += e.solve()
	}

	result := sum
	return result
}

// numbers is map from sorted string pattern to number
var numbers map[string]int = map[string]int{
	"abcefg":  0,
	"cf":      1,
	"acdeg":   2,
	"acdfg":   3,
	"bcdf":    4,
	"abdfg":   5,
	"abdefg":  6,
	"acf":     7,
	"abcdefg": 8,
	"abcdfg":  9,
}

type entry struct {
	patterns []string
	output   []string
}

func (e *entry) solve() int {
	_1 := e.patternByLength(2)[0]
	_7 := e.patternByLength(3)[0]
	_4 := e.patternByLength(4)[0]
	_8 := e.patternByLength(7)[0]

	_2_3_5 := e.patternByLength(5)
	_0_6_9 := e.patternByLength(6)

	_9 := searchContaining(_0_6_9, _4)[0]

	// --> A == 7 - 1
	A := minus(_7, _1)

	// --> G == 9 - 4+A
	G := minus(_9, _4+string(A))

	// --> D == 3 - 1+A+G
	_3_minus_D := _1 + string(A) + string(G)
	_3 := searchContaining(_2_3_5, _3_minus_D)[0]
	D := minus(_3, _3_minus_D)

	// --> B == 9 - 3
	B := minus(_9, _3)

	// --> E == 8 - 9
	E := minus(_8, _9)

	// --> C == 8 - 6
	_6_9 := searchContaining(_0_6_9, string(D))
	var _6 string
	for _, s := range _6_9 {
		if s != _9 {
			_6 = s
			break
		}
	}
	C := minus(_8, _6)

	// --> F == 1 - C
	F := minus(_1, string(C))

	mapping := map[rune]rune{
		A: 'a',
		B: 'b',
		C: 'c',
		D: 'd',
		E: 'e',
		F: 'f',
		G: 'g',
	}

	fmt.Println("A:", string(A), "B:", string(B), "C:", string(C), "D:", string(D), "E:", string(E), "F:", string(F), "G:", string(G))
	fmt.Println("mapping:", mapping)

	result := 0
	for _, o := range e.output {
		chars := strings.Split(o, "")
		for i := 0; i < len(chars); i++ {
			chars[i] = string(mapping[rune(chars[i][0])])
		}
		sort.Strings(chars)
		sorted := strings.Join(chars, "")
		fmt.Println("o:", o, "->", sorted)
		result = result*10 + numbers[sorted]
	}

	fmt.Println("  -->", result)

	return result
}

// patternByLength finds first pattern by length. Supposed to be called for unique lengths.
func (e *entry) patternByLength(l int) []string {
	result := []string{}
	for _, p := range e.patterns {
		if len(p) == l {
			result = append(result, p)
		}
	}
	return result
}

func searchContaining(ss []string, needle string) []string {
	result := []string{}

	for _, s := range ss {
		if contains(s, needle) {
			result = append(result, s)
		}
	}

	return result
}

func contains(heystack, needle string) bool {
outer:
	for _, n := range needle {
		for _, h := range heystack {
			if h == n {
				continue outer
			}
		}
		return false
	}
	return true
}

// minus returns the first character that exists in 's' and not in 't.
// It is supposed to be called in situations where the expected difference is exactly 1 char.
func minus(s, t string) rune {
outer:
	for _, sc := range s {
		for _, tc := range t {
			if tc == sc {
				continue outer
			}
		}
		return sc
	}
	return -1
}
