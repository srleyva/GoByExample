package main

import "fmt"

func sum(nums ...int) {
	fmt.Println(nums, " ")
	total := 0
	for _, num := range nums {
		total = num + total
	}
	fmt.Println(total)
}

func main() {
	sum(1,2,3,4,5,6,7,8)

	nums := []int{1,2,3,4}
	sum(nums...)
}
