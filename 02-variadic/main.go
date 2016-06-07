package main

import "fmt"

func main() {
	greatestNumber := greatest(3, 4, 5, 1, 2, 7)

	fmt.Println(greatestNumber)
}

func greatest(args ...int) int {
	great := 0

	for _, value := range args {
		if value > great {
			great = value
		}
	}

	return great
}
