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
		for x, block := range line {
			next[y][x] = block //default value
			if block == '.' {
				continue
			}
			if !sideOkay(x, y-1, block, lines) {
				continue
			}
			if !sideOkay(x, y+1, block, lines) {
				continue
			}
			if !sideOkay(x-1, y, block, lines) {
				continue
			}
			if !sideOkay(x+1, y, block, lines) {
				continue
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
		return true
	}
	if x < 0 || x >= len(lines[0]) {
		return true
	}
	if lines[y][x] == '.' {
		return block == '#'
	}
	return block == lines[y][x]
}
