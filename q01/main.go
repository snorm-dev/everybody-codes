package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	input, err := os.ReadFile("./input.txt")
	if err != nil {
		log.Fatal(err)
	}
	var potions int
	for _, c := range input {
		switch c {
		case 'A':
		case 'B':
			potions += 1
		case 'C':
			potions += 3
		case '\n':
		default:
			log.Fatalf("invalid enemy: '%s'", string(c))
		}
	}
	fmt.Println(potions)
}
