package main

import (
	"fmt"
	"COURSERA-Functions-Methods-and-Interfaces-in-Go/week1"
	"COURSERA-Functions-Methods-and-Interfaces-in-Go/week2"
)

func main() {

	//WEEK1
	week1.BubbleSort([]int {4,6,999,1,24,2,3,6,8,222})

	//WEEK2
	fmt.Print("Hi user, enter the quantities to calculate the formula: \n")

	fmt.Print("Please enter the acceleration: \n")
	var acceleration float64
    fmt.Scanln(&acceleration)

	fmt.Print("Please enter the initial velocity: \n")
	var initialVelocity float64
    fmt.Scanln(&initialVelocity)

	fmt.Print("Please enter the displacement: \n")
	var displacement float64
    fmt.Scanln(&displacement)

	returnedFunction := week2.GenDisplaceFn(acceleration, initialVelocity, displacement)


	fmt.Print("Please enter also the time: \n")
	var time float64
    fmt.Scanln(&time)

	fmt.Printf("The total displacement after %f is: %f" , time, returnedFunction(time))

}

