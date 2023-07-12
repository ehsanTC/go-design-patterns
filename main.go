package main

import (
	"design-patterns/chain_of_responsibility"
	"design-patterns/observer"
	"fmt"
)

func main() {

	patterns := make([]IPattern, 0)
	patterns = append(patterns, &observer.Observer{Number: 1})
	patterns = append(patterns, &chain_of_responsibility.ATM{Number: 2})

	fmt.Println("\nThe available patterns are:")
	for _, pattern := range patterns {
		fmt.Println(fmt.Sprintf(" %d) %s", pattern.GetNumber(), pattern.Name()))
	}

	fmt.Print("\nEnter pattern a number: ")

	var choice int
	_, err := fmt.Scan(&choice)

	if err != nil || choice > len(patterns) {
		fmt.Println("The entered number is wrong")
		return
	}

	patterns[choice-1].Work()
}
