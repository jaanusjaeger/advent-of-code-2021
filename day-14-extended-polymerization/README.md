# Day 14

https://adventofcode.com/2021/day/14

## Solutions

### Simple / naive

1. Start with the input string
1. On each step apply each extension rule and calculate output string

**Problem:** it quickly gets out of hands because of exponential growth - around at step 25 it becomes unbearable.

### Only count

Principles:
1. We don't need the resulting string, only the counts of character occurrences
1. Given a character pair _pq_ and step count _c_:
   1. the resulting substring (nor the counts of characters) produced by it does not affect the next pair
   2. if there exists an extension rule _pq -> x_, we can apply the same count algorithm for pairs _px_ and _xq_ and merge the resulting counts.

**Problem:** still slow!

### Count with caching

Principles:
1. The variation of input (pair and step count/depth) is quite low
2. The function to calculate character counts return map for the input pair (and step)
3. Apply caching - calculate step key from input pair and depth and look it up from the cache; if not found, add the function result to the cache

**Profit!**
