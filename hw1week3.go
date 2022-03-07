package main

import (
	"fmt"
	"bufio"
	"os"
	"strings"
)

type Animal struct {
	Food string
	Locomotion string
	Noise string
}

func main() {


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
			cow := new(Animal)
			cow.Food = "grass \n"
			cow.Locomotion = "walk \n"
			cow.Noise = "muuuu \n"

			if action == actions[0] {
				//eat
				Eat(cow)
			}else if action == actions[1] {
				//move
				Move(cow)
			}else if action == actions[2] {
				//speak
				Speak(cow)
			} else {
				fmt.Print("Invalid second argument !! \n")
			}
		} else if typeName == animals[1] {
			//bird case
			bird := new(Animal)
			bird.Food = "worms \n"
			bird.Locomotion = "but \n"
			bird.Noise = "cipciop \n"

			if action == actions[0] {
				//eat
				Eat(bird)
			}else if action == actions[1] {
				//move
				Move(bird)
			}else if action == actions[2] {
				//speak
				Speak(bird)
			} else {
				fmt.Print("Invalid second argument !! \n")
			}
		} else if typeName == animals[2] {
			//snake case

			snake := new(Animal)
			snake.Food = "all \n"
			snake.Locomotion = "body \n"
			snake.Noise = "shhhh \n"

			if action == actions[0] {
				//eat
				Eat(snake)
			}else if action == actions[1] {
				//move
				Move(snake)
			}else if action == actions[2] {
				//speak
				Speak(snake)
			} else {
				fmt.Print("Invalid second argument !! \n")
			}
		} else {
			fmt.Print("Invalid first argument !! \n")
		}



	}

}

func Eat(animal *Animal) {
	fmt.Print(animal.Food)
}

func Move(animal *Animal) {
	fmt.Print(animal.Locomotion)
}

func Speak(animal *Animal) {
	fmt.Print(animal.Noise)
}

