package main

import (
	"fmt"
	"sort"
	"strings"
)

var openers string = "([{<"
var closers map[rune]rune

func main() {
	file := openFileFromArgs()
	defer file.Close()

	closers = make(map[rune]rune)
	closers[')'] = '('
	closers[']'] = '['
	closers['}'] = '{'
	closers['>'] = '<'

	lines := []string{}
	scanner := func(line string) error {
		lines = append(lines, line)
		return nil
	}
	scanLines(file, scanner)

	fmt.Println("INPUT:")
	//fmt.Println(lines)

	result1, correct := puzzle1(lines)
	fmt.Println("CORRECT:")
	fmt.Println(correct)
	result2 := puzzle2(correct)

	fmt.Println("RESULT1:", result1)
	fmt.Println("RESULT2:", result2)
}

func puzzle1(lines []string) (int, []string) {
	result := 0
	correct := []string{}

	var scores map[rune]int
	scores = make(map[rune]int)
	scores[')'] = 3
	scores[']'] = 57
	scores['}'] = 1197
	scores['>'] = 25137

	for _, l := range lines {
		s := &stack{}
		valid := true
		for _, r := range l {
			//fmt.Println("----", string(r))
			//fmt.Println(s.toString())
			if strings.IndexRune(openers, r) >= 0 {
				s.push(r)
			} else {
				c, ok := closers[r]
				if !ok {
					result += scores[r]
					valid = false
					break
				}
				p := s.pop()
				if c != p {
					result += scores[r]
					valid = false
					break
				}
			}
		}
		if valid {
			correct = append(correct, l)
		}
	}

	return result, correct
}

func puzzle2(lines []string) int {

	var scores map[rune]int
	scores = make(map[rune]int)
	scores['('] = 1
	scores['['] = 2
	scores['{'] = 3
	scores['<'] = 4

	sc := []int{}

	for _, l := range lines {
		s := &stack{}
		for _, r := range l {
			if strings.IndexRune(openers, r) >= 0 {
				s.push(r)
			} else {
				s.pop()
			}
		}
		fmt.Println("----", l)
		fmt.Println("remaining:", s.toString())
		score := 0
		for r := s.pop(); r != 'X'; r = s.pop() {
			score = score*5 + scores[r]
		}
		fmt.Println("-->", score)
		sc = append(sc, score)
	}

	sort.Ints(sc)

	result := sc[len(sc)/2]

	return result
}
