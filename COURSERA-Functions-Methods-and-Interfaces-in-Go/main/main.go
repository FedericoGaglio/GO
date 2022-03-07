package main

import (
	"fmt"
	"bufio"
	"os"
	"strings"
	//"COURSERA-Functions-Methods-and-Interfaces-in-Go/week1"
	//"COURSERA-Functions-Methods-and-Interfaces-in-Go/week2"
	"COURSERA-Functions-Methods-and-Interfaces-in-Go/week3"
)

func main() {

	//WEEK1
	//week1.BubbleSort([]int {4,6,999,1,24,2,3,6,8,222})

	//WEEK2
	/*fmt.Print("Hi user, enter the quantities to calculate the formula: \n")

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

	fmt.Printf("The total displacement after %f is: %f" , time, returnedFunction(time))*/

	
	//WEEK3
	var animals = []string{"cow","bird","snake"}
	var actions = []string{"eat","move","speak"}

	enterRequest := true

	for enterRequest {

		fmt.Print("Hello user, enter a request please: \n")

		fmt.Print("Remember that the request must have the following pattern: [animal] [action] \n")

		fmt.Print("For the animals you can choose between cow, bird or snake; and for the action you can choose between eat, move or speak \n")

		fmt.Print("Ok, let's go! Enter you request: \n")

		fmt.Print("> \n")


		scanner := bufio.NewScanner(os.Stdin)
		scanner.Scan()
		request := scanner.Text()

		var userInput = make([]string, 2)

		userInput[0] = strings.Split(request, " ")[0]
		userInput[1] = strings.Split(request, " ")[1]

		if len(userInput) != 2 {
			fmt.Print("INVALID INPUT !!! \n")
		}

		typeName := userInput[0]
		action := userInput[1]


		if typeName == animals[0]{
			//cow case
			cow := new(week3.Animal)
			cow.Food = "grass \n"
			cow.Locomotion = "walk \n"
			cow.Noise = "muuuu \n"

			if action == actions[0] {
				//eat
				week3.Eat(cow)
			}else if action == actions[1] {
				//move
				week3.Move(cow)
			}else if action == actions[2] {
				//speak
				week3.Speak(cow)
			} else {
				fmt.Print("Invalid second argument !! \n")
			}
		} else if typeName == animals[1] {
			//bird case
			bird := new(week3.Animal)
			bird.Food = "worms \n"
			bird.Locomotion = "but \n"
			bird.Noise = "cipciop \n"

			if action == actions[0] {
				//eat
				week3.Eat(bird)
			}else if action == actions[1] {
				//move
				week3.Move(bird)
			}else if action == actions[2] {
				//speak
				week3.Speak(bird)
			} else {
				fmt.Print("Invalid second argument !! \n")
			}
		} else if typeName == animals[2] {
			//snake case

			snake := new(week3.Animal)
			snake.Food = "all \n"
			snake.Locomotion = "body \n"
			snake.Noise = "shhhh \n"

			if action == actions[0] {
				//eat
				week3.Eat(snake)
			}else if action == actions[1] {
				//move
				week3.Move(snake)
			}else if action == actions[2] {
				//speak
				week3.Speak(snake)
			} else {
				fmt.Print("Invalid second argument !! \n")
			}
		} else {
			fmt.Print("Invalid first argument !! \n")
		}



	}

}

