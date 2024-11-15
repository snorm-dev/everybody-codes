package main

import (
	"bytes"
	"fmt"
	"log"
	"os"
	"regexp"
)

func main() {
	input, err := os.ReadFile("./input.txt")
	if err != nil {
		log.Fatal(err)
	}
	r := regexp.MustCompile("(?s)^WORDS:(.*)\n\n(.*)\n$")
	matches := r.FindSubmatch(input)
	words := bytes.Split(matches[1], []byte{','})
	wordMap := make(map[string]bool)
	for _, word := range words {
		rev := make([]byte, len(word))
		for i, c := range word {
			rev[len(rev)-1-i] = c
		}
		wordMap[string(word)] = true
		wordMap[string(rev)] = true
	}

	lines := bytes.Split(matches[2], []byte{'\n'})

	count := 0
	for _, line := range lines {
		next := 0
		for i := range line {
			for word := range wordMap {
				if i+len(word) > len(line) {
					continue
				}
				poss := line[i : i+len(word)]
				if word == string(poss) {
					overlap := max(0, next-i)
					count += max(0, len(word)-overlap)

					next = max(next, i+len(word))
				}
			}
		}
	}
	fmt.Println(count)
}
