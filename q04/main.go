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

	heights := make([]int, len(lines))
	for i, line := range lines {
		heights[i], err = strconv.Atoi(string(line))
		if err != nil {
			log.Fatal("invalid int in input!?")
		}
	}

	slices.Sort(heights)

	median := heights[len(heights)/2]

	count := 0
	for _, height := range heights {
		diff := height - median
		if diff < 0 {
			diff *= -1
		}
		count += diff
	}
	fmt.Println(count)
}
