package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func openFileFromArgs() *os.File {
	fname := os.Args[1]
	fmt.Println("Arguments:", os.Args)

	file, err := os.Open(fname)
	if err != nil {
		log.Fatal(err)
	}

	return file
}

func scanLines(file *os.File, callback func(string) error) {
	scanner := bufio.NewScanner(file)
	// optionally, resize scanner's capacity for lines over 64K, see next example
	for scanner.Scan() {
		line := scanner.Text()
		fmt.Println(line)
		if len(line) == 0 {
			continue
		}

		if err := callback(line); err != nil {
			log.Fatal(err, "Error from callback")
		}
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}

func parseInt(s string) int {
	i, err := strconv.Atoi(s)
	if err != nil {
		log.Fatal(err, "Invalid input", s)
	}
	return i
}

func intFromBinaryString(s string) int {
	if n, err := strconv.ParseInt(s, 2, 64); err != nil {
		log.Fatal(err, "Invalid input", s)
		return 0
	} else {
		return int(n)
	}
}
