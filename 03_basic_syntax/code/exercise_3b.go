package main

import (
	"fmt"
	"strings"
)

func main() {

	var name, city, niceWeatherText string
	var yearsLived int
	var isNicWeather bool

	fmt.Println("Please enter your name")
	fmt.Scan(&name)

	fmt.Println("What city do you live?")
	fmt.Scan(&city)

	fmt.Println("How long have you lived there?")
	fmt.Scan(&yearsLived)

	fmt.Println("Is the weather nice there?")
	fmt.Scan(&niceWeatherText)

	lowerCasedWeatherAnswer := strings.ToLower(niceWeatherText)
	if lowerCasedWeatherAnswer == "yes" {
		isNicWeather = true
	} else {
		isNicWeather = false
	}

	fmt.Printf("Hi! My name is %s. I have lived in %s for %d years. They say the weather is amazing, which is %t", name, city, yearsLived, isNicWeather)

}
