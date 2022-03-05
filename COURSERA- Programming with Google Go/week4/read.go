package week4

import (
	"fmt"
	"os"
	"log"
	"bufio"
	"strings"
)

type Person struct {
	Fname string
	Lname string
}

func InsertData() {


	lines, err := os.Open("/Users/federicogaglio/Desktop/GO/COURSERA- Programming with Google Go/week4/names.txt")

	if(err!=nil) {
		log.Fatal(err)
	}

	defer lines.Close()
	scanner := bufio.NewScanner(lines)

	count := 0

	var structList = make([]Person, count) 

	for scanner.Scan() {
        count += 1

		currentLine := scanner.Text()

		wordsOnLine := strings.Fields(currentLine) 

		person := Person{Fname: wordsOnLine[0], Lname: wordsOnLine[1]}
		structList = append(structList, person)
	}

	fmt.Print(structList)
}
