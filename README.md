# Advent of Code 2021

https://adventofcode.com/2021

Each day is in a separate subdirectory.

Subdirectory layout:

    ├── Makefile           # The makefile
    ├── example-input      # Test input from the puzzle text
    ├── go.mod             # The Go module file
    ├── input              # Real input
    ├── main.go            # The executable
    ├── <other>.go         # Possible helpers and utils
    └── puzzle             # Binary (build output)


## Usage

    # Run with test input
    make test

    # Run with real input
    make run
