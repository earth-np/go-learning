package main

import "fmt"

func vals(nums ...int) []int {
	return nums
}

func main() {
	a := []int{0, 1, 2, 3, 4}
	res := vals(a...)
	fmt.Println((res))

}
