package main

import "fmt"

func vals() (int, int, int, string) {
	return 3, 7, 3, "earth"
}

func main() {

	a, _, _, e := vals()
	fmt.Println(a)
	fmt.Println(e)

}
