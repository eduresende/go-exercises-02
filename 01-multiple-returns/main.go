package main

import "fmt"

// func myfunc(x int) (int, bool) {
// 	first := x / 2
// 	second := (x%2 == 0)

// 	return first, second
// }

func main() {
	fmt.Println(
		func(x int) (int, bool) {
			first := x / 2
			second := (x%2 == 0)

			return first, second
		}(4))
}
