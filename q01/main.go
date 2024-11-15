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
	input = input[:len(input)-1]
	if len(input)%2 != 0 {
		log.Fatal("input has odd number of enemies")
	}

	var potions int
	i := 0
	for i < len(input) {
		potions += calculatePotionsNeededForPair(input[i], input[i+1])
		i += 2
	}
	fmt.Println(potions)
}

func calculatePotionsNeededForPair(a, b byte) int {
	return max(0, 2+calculatePotionsNeededForEnemy(a)+calculatePotionsNeededForEnemy(b))
}

func calculatePotionsNeededForEnemy(c byte) int {
	switch c {
	case 'A':
		return 0
	case 'B':
		return 1
	case 'C':
		return 3
	case 'D':
		return 5
	case 'x':
		return -2
	default:
		log.Fatalf("invalid enemy: '%s'", string(c))
	}
	panic("unreachable")
}
