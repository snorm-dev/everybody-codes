package main

import (
	"bytes"
	"fmt"
	"log"
	"math"
	"os"
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

	minHeight := math.MaxInt
	for _, height := range heights {
		minHeight = min(minHeight, height)
	}

	count := 0
	for _, height := range heights {
		count += height - minHeight
	}
	fmt.Println(count)
}
