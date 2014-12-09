package main

import "fmt"

func main() {
	fmt.Print("Enter the current temperature (Fahrenheit): ")
	var temp float64
	fmt.Scanf("%f", &temp)

	output := (temp - 32) * 5/9

	fmt.Println("Current temperature is", output)
}