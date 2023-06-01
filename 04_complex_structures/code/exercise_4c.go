package main

import (
	"fmt"
)

func main() {
	// Part 1
	scores := []float64{1.53, 32.632123, 5.43, 23.6, 21.78}
	fmt.Println(average(scores...))

	// Part 2
	pets := map[string]string{
		"joey":    "dog",
		"kitty":   "cat",
		"jerry":   "mouse",
		"tom":     "cat",
		"tornado": "horse",
		"zombe":   "dog",
	}
	hasPet(pets, "ann")

	// Part 3
	groceries := []string{"milk", "yogourt"}
	newGroceries := []string{"haa", "hoooo"}
	addGrocery(groceries, newGroceries...)
}

// Part 1
func average(nums ...float64) float64 {
	var sum float64 = 0
	for _, num := range nums {
		sum += num
	}
	return sum / float64(len(nums))
}

// Part 2
func hasPet(pets map[string]string, name string) bool {
	if _, ok := pets[name]; ok {
		fmt.Println(name, "exists")
		return ok
	} else {
		fmt.Println(name, "does NOT exist")
		return ok
	}
}

// Part 3
func addGrocery(groceries []string, items ...string) {
	newGroceries := append(groceries, items...)
	fmt.Println(newGroceries)
}
