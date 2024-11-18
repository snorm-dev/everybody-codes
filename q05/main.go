package main

import (
	"bytes"
	"fmt"
	"log"
	"os"
	"slices"
	"strconv"
	"strings"
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

	seen := make(map[string]bool)
	maxFront := ""

	for i := 0; ; i++ {
		col := i % len(columns)
		clapper := columns[col][0]
		columns[col] = columns[col][1:]

		nextCol := (i + 1) % len(columns)
		columns[nextCol] = round(clapper, columns[nextCol])

		front := strings.Builder{}
		for _, column := range columns {
			front.WriteString(strconv.Itoa(column[0]))
		}
		maxFront = max(maxFront, front.String())

		sb := strings.Builder{}
		for _, column := range columns {
			for _, person := range column {
				sb.WriteString(strconv.Itoa(person))
				sb.WriteByte(';')
			}
			sb.WriteByte('|')
		}
		snapshot := sb.String()

		if seen[snapshot] {
			fmt.Println(maxFront)
			return
		}
		seen[snapshot] = true
	}
}

func round(clapper int, line []int) []int {
	index := absorptionIndex(clapper, len(line))
	return slices.Insert(line, index, clapper)
}

func absorptionIndex(clapper, length int) int {
	clapper = 1 + (clapper-1)%(2*length)

	if clapper > length {
		return 2*length + 1 - clapper
	}
	return clapper - 1
}
