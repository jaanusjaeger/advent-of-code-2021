# Day 20

https://adventofcode.com/2021/day/20

## Solutions

Observations:
* The infinite edge could result in infinite number of lit pixels, but
* The enhancement algorithm turns 3x3 of '.' into '#' and 3x3 of '#' into '.'
* Hence, after even number of steps all "infinite" edges of '#' will be '.' again
* Add sufficient (finite) buffer around the initial matrix, to preserve all non-infinite enhancements.

For puzzle 2:
* Simply add bigger buffer around the matrix
