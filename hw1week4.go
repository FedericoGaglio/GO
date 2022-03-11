package main

import (
	"fmt"
	"bufio"
	"os"
	"strings"
	"log"
)


type Animal interface {
	Eat()
	Move()
	Speak()
	GetName() string
}

type Cow struct{ Name string }
type Snake struct{ Name string }
type Bird struct{ Name string }

func (c Cow) GetName() string {
	return c.Name
}

func (s Snake) GetName() string {
	return s.Name
}

func (b Bird) GetName() string {
	return b.Name
}

func (c Cow) Eat() {
	fmt.Printf("%s Eats grass\n", c.Name)
}

func (c Cow) Move() {
	fmt.Printf("%s walks\n", c.Name)
}

func (c Cow) Speak() {
	fmt.Printf("%s moos\n", c.Name)
}

func (s Snake) Eat() {
	fmt.Printf("%s Eats mice\n", s.Name)
}

func (s Snake) Move() {
	fmt.Printf("%s slithers\n", s.Name)
}

func (s Snake) Speak() {
	fmt.Printf("%s hisses\n", s.Name)
}

func (b Bird) Eat() {
	fmt.Printf("%s Eats worms\n", b.Name)
}

func (b Bird) Move() {
	fmt.Printf("%s flys\n", b.Name)
}

func (b Bird) Speak() {
	fmt.Printf("%s peeps and chirps\n", b.Name)
}


func main() {


	inputReader := bufio.NewReader(os.Stdin)
	sliceOfAnimals := []Animal{}

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
				cow := new(Cow)
				cow.Name = sliceOfQuery[1]
				sliceOfAnimals = append(sliceOfAnimals, cow)
				fmt.Println("Created it!")
			} else if sliceOfQuery[2] == "snake" {
				snake := new(Snake)
				snake.Name = sliceOfQuery[1]
				sliceOfAnimals = append(sliceOfAnimals, snake)
				fmt.Println("Created it!")
			} else if sliceOfQuery[2] == "bird" {
				bird := new(Bird)
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