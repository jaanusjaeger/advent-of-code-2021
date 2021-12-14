package main

import (
	"fmt"
	"strings"
)

var cache map[string](map[byte]int64)

func main() {
	file := openFileFromArgs()
	defer file.Close()

	template := ""
	rules := map[string]string{}
	scanner := func(line string) error {
		if template == "" {
			template = line
			return nil
		}
		ss := strings.Split(line, " -> ")
		if len(ss) != 2 {
			return fmt.Errorf("invalid input: %s", line)
		}
		rules[ss[0]] = ss[1]
		return nil
	}
	scanLines(file, scanner)

	fmt.Println("TEMPLATE:")
	fmt.Println(template)
	fmt.Println("RULES:")
	fmt.Println(rules)

	cache = map[string](map[byte]int64){}

	result1 := puzzle1(template, rules, 10)
	result2 := puzzle2(template, rules, 40)

	fmt.Println("RESULT1:", result1)
	fmt.Println("RESULT2:", result2)
}

func puzzle1(input string, rules map[string]string, steps int) string {
	result := input

	fmt.Println("STEP", 0, ":", result)
	for i := 0; i < steps; i++ {
		result = step(result, rules)
		//fmt.Println("STEP", i+1, ":", result)
	}

	counts := map[byte]int{}
	for _, c := range result {
		counts[byte(c)] += 1
	}
	fmt.Println("COUNTS:", counts)
	max := -1
	min := -1
	for _, count := range counts {
		if max == -1 || count > max {
			max = count
		}
		if min == -1 || count < min {
			min = count
		}
	}

	return fmt.Sprintf("%d", max-min)
}

func puzzle2(input string, rules map[string]string, steps int) string {
	counts := map[byte]int64{}
	window := ""
	for i := range input {
		counts[input[i]]++
		if i < 1 {
			continue
		}
		window = input[i-1 : i+1]
		c := countOccurrences(window, rules, steps)
		merge(counts, c)
	}

	fmt.Println("COUNTS:", counts)
	max := int64(-1)
	min := int64(-1)
	for _, count := range counts {
		if max == -1 || count > max {
			max = count
		}
		if min == -1 || count < min {
			min = count
		}
	}

	return fmt.Sprintf("%d", max-min)
}

func step(input string, rules map[string]string) string {
	result := []byte{}

	window := ""
	for i := range input {
		if i < 1 {
			result = append(result, input[i])
			continue
		}
		window = input[i-1 : i+1]
		if r, ok := rules[window]; ok {
			result = append(result, []byte(r)...)
		}
		result = append(result, window[1])
	}

	return string(result)
}

func countOccurrences(in string, rules map[string]string, countdown int) map[byte]int64 {
	key := fmt.Sprintf("%s/%d", in, countdown)
	if count, ok := cache[key]; ok {
		return count
	}

	counts := map[byte]int64{}
	if countdown == 0 {
		return counts
	}
	if countdown == 40 {
		fmt.Println("  step:", in)
	}

	r, ok := rules[string(in)]
	if !ok {
		return counts
	}
	counts[r[0]]++
	c1 := countOccurrences(string([]byte{in[0], r[0]}), rules, countdown-1)
	c2 := countOccurrences(string([]byte{r[0], in[1]}), rules, countdown-1)
	merge(counts, c1)
	merge(counts, c2)
	cache[key] = counts
	return counts
}

func merge(target, source map[byte]int64) {
	for b, c := range source {
		target[b] += c
	}
}
