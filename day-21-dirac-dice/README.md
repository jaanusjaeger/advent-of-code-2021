# Day 21

https://adventofcode.com/2021/day/21

## Solutions

For puzzle 2:
* The growth of data is enormous - don't try to keep every universe in memory
* Instead, observe that the number of combinations of possible game statuses
  (position and score for both players) is rather limited (`pos=1..10`, `score=0..2?`)
* Use map from game status into number of universes it is used in
