package main

import "fmt"

// fibonacci is a function that returns
// a function that returns an int.
func fibonacci() func() int {
	nums := []int {0, 1, 1}

	return func() int {
		cur := nums[0]
		
		nums[0] = nums[1]
		nums[1] = nums[2]
		nums[2] = nums[0] + nums[1]
		
		return cur
	}
}

func main() {
	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}
}
