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
	r := regexp.MustCompile("WORDS:(.*)\n\n(.*)\n")
	matches := r.FindSubmatch(input)
	words := bytes.Split(matches[1], []byte{','})
	count := 0
	for _, word := range words {
		text := matches[2]
		i := bytes.Index(text, word)
		for i >= 0 {
			count += 1
			text = text[i+len(word):]
			i = bytes.Index(text, word)
		}
	}
	fmt.Println(count)
}
