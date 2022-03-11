package week4

import (
	"fmt"
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