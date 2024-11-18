package main

import (
	"bytes"
	"fmt"
	"log"
	"os"
)

func main() {
	input, err := os.ReadFile("./input.txt")
	if err != nil {
		log.Fatal(err)
	}
	input = bytes.TrimRight(input, "\n")

	lines := bytes.Split(input, []byte{'\n'})

	count := 0
	for i, lines := step(lines); i > 0; i, lines = step(lines) {
		count += i
	}
	fmt.Println(count)
}

func step(lines [][]byte) (int, [][]byte) {

	next := make([][]byte, len(lines))
	for i := range next {
		next[i] = make([]byte, len(lines[0]))
	}

	count := 0
	for y, line := range lines {
	blockLoop:
		for x, block := range line {
			next[y][x] = block //default value
			if block == '.' {
				continue
			}
			adjacentBlocks := [][]int{
				{x, y - 1},
				{x, y + 1},
				{x - 1, y},
				{x + 1, y},
				{x - 1, y - 1},
				{x + 1, y + 1},
				{x - 1, y + 1},
				{x + 1, y - 1},
			}
			for _, adj := range adjacentBlocks {
				if !sideOkay(adj[0], adj[1], block, lines) {
					continue blockLoop
				}
			}

			if block+1 == 0 {
				panic("too deep: overflow!")
			}
			if block == '#' {
				next[y][x] = '1'
			} else {
				next[y][x] = block + 1
			}
			count++
		}
	}

	return count, next
}

func sideOkay(x, y int, block byte, lines [][]byte) bool {
	if y < 0 || y >= len(lines) {
		return block == '#'
	}
	if x < 0 || x >= len(lines[0]) {
		return block == '#'
	}
	if lines[y][x] == '.' {
		return block == '#'
	}
	return block == lines[y][x]
}
