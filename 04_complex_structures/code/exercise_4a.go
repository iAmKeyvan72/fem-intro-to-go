package main

import "fmt"

func average(nums ...float64) float64 {
	sum := 0.0

	for _, num := range nums {
		sum += num
	}

	return sum / float64(len(nums))
}

func main() {
	fmt.Println(average(42, 11, 49))
}
