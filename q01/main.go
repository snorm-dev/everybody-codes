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
	if len(input)%3 != 0 {
		log.Fatal("input has odd number of enemies")
	}

	var potions int
	i := 0
	for i < len(input) {
		potions += calculatePotionsNeededForBattle(input[i : i+3])
		i += 3
	}
	fmt.Println(potions)
}

func calculatePotionsNeededForBattle(enemies []byte) int {
	numEnemies := len(enemies)
	potions := 0
	for _, e := range enemies {
		p, ok := calculatePotionsNeededForEnemy(e)
		if !ok {
			numEnemies -= 1
			continue
		}
		potions += p
	}
	potions += numEnemies * (numEnemies - 1)

	return max(0, potions)
}

func calculatePotionsNeededForEnemy(c byte) (int, bool) {
	switch c {
	case 'A':
		return 0, true
	case 'B':
		return 1, true
	case 'C':
		return 3, true
	case 'D':
		return 5, true
	case 'x':
		return 0, false
	default:
		log.Fatalf("invalid enemy: '%s'", string(c))
	}
	panic("unreachable")
}
