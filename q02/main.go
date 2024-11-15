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
	lines := bytes.Split(matches[2], []byte{'\n'})

	isRune := make([][]bool, len(lines))
	for i := range isRune {
		isRune[i] = make([]bool, len(lines[0]))
	}
	for y, line := range lines {
		for x := range line {
			for _, word := range words {
				poss := make([]byte, len(word))
				// right
				{
					for i := range poss {
						off := wrapMod(x+i, len(line))
						poss[i] = lines[y][off]
					}
					if bytes.Equal(word, poss) {
						for i := range poss {
							off := wrapMod(x+i, len(line))
							isRune[y][off] = true
						}
					}
				}
				// left
				{
					for i := range poss {
						off := wrapMod(x-i, len(line))
						poss[i] = lines[y][off]
					}
					if bytes.Equal(word, poss) {
						for i := range poss {
							off := wrapMod(x-i, len(line))
							isRune[y][off] = true
						}
					}
				}
				// down
				if y+len(poss)-1 < len(lines) {
					for i := range poss {
						off := y + i
						poss[i] = lines[off][x]
					}
					if bytes.Equal(word, poss) {
						for i := range poss {
							off := y + i
							isRune[off][x] = true
						}
					}
				}
				// up
				if y-len(poss)+1 >= 0 {
					for i := range poss {
						off := y - i
						poss[i] = lines[off][x]
					}
					if bytes.Equal(word, poss) {
						for i := range poss {
							off := y - i
							isRune[off][x] = true
						}
					}
				}
			}
		}
	}
	count := 0
	for y := range isRune {
		for x := range isRune[y] {
			if isRune[y][x] {
				count++
			}
		}
	}
	fmt.Println(count)
}

func wrapMod(x, n int) int {
	r := x % n // if x <= 0, then -n < r <= 0
	r += n
	r %= n
	return r
}
