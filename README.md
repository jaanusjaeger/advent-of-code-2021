# Advent of Code 2021

https://adventofcode.com/2021

Solution template is in the `template/` directory and contains the most up to date solution structure and superset of all helper files (`functions.g`, `matrix.go`, etc).

Solution for each day is in a separate subdirectory. This may contain subset or modified versions of the helper files.

Directory layout:

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
