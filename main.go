package main

import (
	"design-patterns/chain_of_responsibility"
	"design-patterns/observer"
	"fmt"
)

func main() {
	println("Enter pattern number: \n",
		"1) Observer\n",
		"2) Chain of responsibility\n")

	var choice int
	_, err := fmt.Scan(&choice)

	if err != nil {
		fmt.Println("The entered number is wrong")
		return
	}

	switch choice {
	case 1:
		chain_of_responsibility.Work()
	case 2:
		observer.Work()
	}
}
