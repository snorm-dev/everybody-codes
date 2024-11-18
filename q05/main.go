package main

import (
	"bytes"
	"fmt"
	"log"
	"os"
	"slices"
	"strconv"
)

func main() {
	input, err := os.ReadFile("./input.txt")
	if err != nil {
		log.Fatal(err)
	}
	input = bytes.TrimRight(input, "\n")
	lines := bytes.Split(input, []byte{'\n'})

	columns := make([][]int, bytes.Count(lines[0], []byte{' '})+1)
	for _, line := range lines {
		nums := bytes.Split(line, []byte{' '})
		for i, s := range nums {
			n, err := strconv.Atoi(string(s))
			if err != nil {
				log.Fatal("bad input")
			}
			columns[i] = append(columns[i], n)
		}
	}

	counts := make(map[int]int)

	for i := 0; ; i++ {
		col := i % len(columns)
		clapper := columns[col][0]
		columns[col] = columns[col][1:]

		nextCol := (i + 1) % len(columns)
		columns[nextCol] = round(clapper, columns[nextCol])

		front := 0
		for _, column := range columns {
			for c := column[0]; c > 0; c /= 10 {
				front *= 10
			}
			front += column[0]
		}
		counts[front]++
		if counts[front] == 2024 {
			fmt.Println((i + 1) * front)
			return
		}
	}
}

func round(clapper int, line []int) []int {
	index := absorptionIndex(clapper, len(line))
	return slices.Insert(line, index, clapper)
}

func absorptionIndex(clapper, length int) int {
	// examples
	// assume they run back around ad infinitum
	// 1 4 -> 0
	// 2 4 -> 1
	// 3 4 -> 2
	// 4 4 -> 3
	// 5 4 -> 4 = 2n + 1 - c
	// 6 4 -> 3
	// 7 4 -> 2
	// 8 4 -> 1
	// 9 4 -> 0 // same as 1
	clapper = clapper % (2 * length)

	if clapper > length {
		return 2*length + 1 - clapper
	}
	return clapper - 1
}
