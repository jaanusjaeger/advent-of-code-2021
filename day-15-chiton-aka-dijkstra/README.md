# Day 15

https://adventofcode.com/2021/day/15

## Solutions

* Dijkstra
* Priority queue implemented by sorted array, using binary search (i.e. not the regular binary heap)
* Hint: in normal graph a _D-tour_ could be a shortcut, depending on the edges' weights, but
  in this case the weight of virtual edge is the weight of the node.

  This simplifies the algorithm by not having to reconsider nodes that are already in the PQ.
  I.e. no need to reorganize the element when it's distance changes.
