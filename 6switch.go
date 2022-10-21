package main

import "fmt"

func main() {
	i := 1
	for i <= 30 {
		fmt.Println(i)
		switch {
		case i%4 == 3:
			fmt.Println("tri")
		case i%4 == 2:
			fmt.Println("di")
		case i%4 == 1:
			fmt.Println("uno")
		default:
			fmt.Println("zero")
		}
		// switch i % 4 {
		// case 3:
		// 	fmt.Println("tri")
		// case 2:
		// 	fmt.Println("di")
		// case 1:
		// 	fmt.Println("uno")
		// default:
		// 	fmt.Println("zero")
		// }
		i++
	}
}
