package main

import "fmt"

func main() {
	fmt.Print("Enter measurement in feet: ")
	var feet float64
	fmt.Scanf("%f", &feet)

	output := feet * 0.3048

	fmt.Println("The length in meters is", output)
}