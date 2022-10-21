package main

import "fmt"

func main() {
	i := 1
	for i <= 30 {
		fmt.Println(i)
		i++
		if i%2 == 0 {
			fmt.Println("Even")
		} else {
			fmt.Println("Odd")

		}
	}
}
