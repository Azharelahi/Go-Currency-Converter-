package main

import (
	"fmt"
	"os"
)

func main() {
	var myChoice string
	var myTarget string
	var target int
	var choice int
	var arr2 []string
	var value float32
	arr := []string{"celsius", "Kelvin", "Fahrenheit"}
	fmt.Println("Enter the source unit from the given list")
	for i := 0; i < len(arr); i++ {
		fmt.Println(i+1, ". ", arr[i])
	}
	_, er := fmt.Scanln(&choice)

	if choice <= 0 || choice > len(arr) {
		fmt.Println("Wrong choice")
		os.Exit(1)
	} else if er != nil {
		fmt.Println("Only integer value is allowed")
		os.Exit(5)
	}
	myChoice = arr[choice-1]
	fmt.Println("Enter the Target unit")
	for i := 0; i < len(arr); i++ {
		if choice != i+1 {
			arr2 = append(arr2, arr[i])
		}
		continue
	}
	for i := 0; i < len(arr2); i++ {
		fmt.Println(i+1, ". ", arr2[i])
	}
	fmt.Scanln(&target)
	if target <= 0 || target > len(arr2) {
		fmt.Println("Wrong input")
		os.Exit(2)
	}
	target = target - 1
	myTarget = arr[target]
	for i := 0; i < len(arr); i++ {
		if target == i {
			fmt.Println("Enter the value of temperature in : ", myChoice)
		}
	}
	_, err := fmt.Scanln(&value)
	if err != nil {
		fmt.Println("Only numbers are allowed to enter")
		os.Exit(3)
	}
	result := printResult(value, myChoice, myTarget)
	fmt.Println("Temperature in ", myTarget, " is ", result)
}
func printResult(val float32, choice string, target string) float32 {
	if choice == "celsius" && target == "Kelvin" {
		return val + 273.15
	} else if choice == "celsius" && target == "Fahrenheit" {
		return (val * 9 / 5) + 32
	} else if choice == "Kelvin" && target == "celsius" {
		return val - 273.15
	} else if choice == "Kelvin" && target == "Fahrenheit" {
		return (val-273.15)*9/5 + 32
	} else if choice == "Fahrenheit" && target == "celsius" {
		return (val - 32) * 5 / 9
	} else if choice == "Fahrenheit" && target == "Kelvin" {
		return (val-32)*5/9 + 273.15
	} else {
		fmt.Println("Invalid conversion type.")
		return 0
	}
}
