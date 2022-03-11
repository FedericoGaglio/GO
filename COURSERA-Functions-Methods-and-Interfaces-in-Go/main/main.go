package main

import (
	"fmt"
	"bufio"
	"os"
	"strings"
	"log"
	//"COURSERA-Functions-Methods-and-Interfaces-in-Go/week1"
	//"COURSERA-Functions-Methods-and-Interfaces-in-Go/week2"
	//"COURSERA-Functions-Methods-and-Interfaces-in-Go/week3"
	"COURSERA-Functions-Methods-and-Interfaces-in-Go/week4"
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
	/*var animals = []string{"cow","bird","snake"}
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



	}*/



	//WEEK4

	inputReader := bufio.NewReader(os.Stdin)
	sliceOfAnimals := []week4.Animal{}

	fmt.Println("Welcome to Animal Park where you can create and query animals")
	fmt.Println("To begin, you can do either of the following types of queries:")
	fmt.Println("To create an animal for the database, you will need to use the following Syntax:")
	fmt.Println("newanimal (name of animal) (type of animal)")
	fmt.Println("Upon creation of an animal, you can then query the animal(s) by doing the following:")
	fmt.Println("query (name of animal) (action of animal)")

	for {

		fmt.Print(">")
		userQuery, err := inputReader.ReadString('\n')
		if err != nil {
			log.Fatal(err)
		}

		//userQuery = strings.Trim(userQuery, "\n")
		userQuery = userQuery[:len(userQuery)-2]
		sliceOfQuery := strings.Split(userQuery, " ")
		if len(sliceOfQuery) > 3 || len(sliceOfQuery) < 3 {
			fmt.Println("Invalid query 1")
		}

		determinant := sliceOfQuery[0]
		switch determinant {
		case "newanimal":

			if sliceOfQuery[2] == "cow" {
				cow := new(week4.Cow)
				cow.Name = sliceOfQuery[1]
				sliceOfAnimals = append(sliceOfAnimals, cow)
				fmt.Println("Created it!")
			} else if sliceOfQuery[2] == "snake" {
				snake := new(week4.Snake)
				snake.Name = sliceOfQuery[1]
				sliceOfAnimals = append(sliceOfAnimals, snake)
				fmt.Println("Created it!")
			} else if sliceOfQuery[2] == "bird" {
				bird := new(week4.Bird)
				bird.Name = sliceOfQuery[1]
				sliceOfAnimals = append(sliceOfAnimals, bird)
				fmt.Println("Created it!")
			} else {
				fmt.Println("Invalid Query 2")
				fmt.Println(sliceOfAnimals)
			}
		case "query":

			var animals = []string{"cow","snake","bird"}

			for _, animal := range sliceOfAnimals {
				for _,a := range animals{
					if a == sliceOfQuery[1] {
						if sliceOfQuery[2] == "move" {
							animal.Move()
						} else if sliceOfQuery[2] == "eat" {
							animal.Eat()
						} else if sliceOfQuery[2] == "speak" {
							animal.Speak()
						}
					}
				}
			}

		default:
			fmt.Println("Invalid user query 3")

		}
	}
}

