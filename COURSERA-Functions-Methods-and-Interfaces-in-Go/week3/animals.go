package week3

import (
	"fmt"
)

type Animal struct {
	Food string
	Locomotion string
	Noise string
}

// funzioni che usano la struct e pertanto risultano essere quindi generiche

func Eat(animal *Animal) {
	fmt.Print(animal.Food)
}

func Move(animal *Animal) {
	fmt.Print(animal.Locomotion)
}

func Speak(animal *Animal) {
	fmt.Print(animal.Noise)
}