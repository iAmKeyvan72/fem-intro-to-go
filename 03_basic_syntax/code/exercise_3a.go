package main

import "fmt"

func main() {
	first := "Hey u"

	for index, c := range first {
		if index%2 == 1 {
			fmt.Println(string(c))
		}
	}
}
