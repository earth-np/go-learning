package main

import "fmt"

func main() {
	numbers := [5]int{1, 0, 4, 2, 3}
	// numbers := [5]int{0, 0, 0, 0, 0}
	i := 0
	for i <= 4 {

		fmt.Println((numbers[i]))
		i++
	}
}
