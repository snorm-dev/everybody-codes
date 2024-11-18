package main

import (
	"bytes"
	"fmt"
	"log"
	"os"
	"slices"
)

func main() {
	input, err := os.ReadFile("./input.txt")
	if err != nil {
		log.Fatal(err)
	}
	input = bytes.TrimRight(input, "\n")
	lines := bytes.Split(input, []byte{'\n'})

	columns := make([][]int, (len(lines[0])+1)/2)
	for _, line := range lines {
		for i, c := range line {
			if c == ' ' {
				continue
			}
			columns[i/2] = append(columns[i/2], int(c-'0'))
		}
	}

	for i := range 10 {
		col := i % len(columns)
		clapper := columns[col][0]
		columns[col] = columns[col][1:]

		nextCol := (i + 1) % len(columns)
		columns[nextCol] = round(clapper, columns[nextCol])

		for _, column := range columns {
			fmt.Print(column[0])
		}
		fmt.Println()
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
