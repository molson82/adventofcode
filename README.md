<p align="center"><img src="aoc20.png"></p>

**2020**

Here lies my solutions to [Advent of Code 2020](https://adventofcode.com/2020), an Advent calendar full of programming puzzles from December 1st all the way to Christmas.
## Implementation Goals

The solutions posted here are NOT cleaned-up versions and I will not be aiming for the leaderboards. For all solutions, the main implementation goals were, in descending order:

* **Readability:** Clean, readable, self-explanatory and commented code above all else.
* **Input Generalization:** Should work not only for my input but for anyone's, with some assumptions made about it, which are noted when appropriate. 
* **Speed:** Use efficient algorithms, keeping runtime reasonably low without extreme micro-optimizations.
* **Minimal Imports:** Refrain from `import`s besides utilities (`strings`, `time`, `fmt`) and basic standard libraries (`math`, `itertools`, `collections`). When the knowledge of functions and structures are considered vital to the problem solution (graphs, trees, linked lists, union-find, etc.), reimplement them.

## Thanks!

Many thanks to [Eric Wastl](http://was.tl/), who creates Advent of Code, as well as to the amazing community over at [/r/adventofcode](https://www.reddit.com/r/adventofcode/)!

```
cd internal/day1
go test -v ./ -coverprofile=./coverage.out -coverpkg ./... -cpuprofile ./cpu.out -memprofile ./mem.out
```