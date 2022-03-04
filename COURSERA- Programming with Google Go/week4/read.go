package week4

import (
	"fmt"
	"os"
	"log"
	"bufio"
	"strings"
	//"encoding/json"
	//"reflect"
)

type Person struct {
	Fname string
	Lname string
}

func InsertData() {


	lines, err := os.Open("/Users/federicogaglio/Desktop/courserago/week4/names.txt")

	if(err!=nil) {
		log.Fatal(err)
	}

	defer lines.Close()
	scanner := bufio.NewScanner(lines)

	count := 0
	//var lineFileMap = make( map[string] string )

	var structList = make([]Person, count) //sara il mio slice di struct

	for scanner.Scan() {
        count += 1

		currentLine := scanner.Text()

		wordsOnLine := strings.Fields(currentLine) //ALTERNATIVA fields := strings.Split(string(currentLine), " ")
		//lineFileMap[wordsOnLine[0]] = wordsOnLine[1]

		person := Person{Fname: wordsOnLine[0], Lname: wordsOnLine[1]}
		structList = append(structList, person)
	}

	fmt.Print(structList)

	//Reflect -> libreria che permette di stampare i TIPI degli oggetti
	//for pos,_ := range structList {
		//fmt.Print(reflect.TypeOf(structList[pos]))
		//fmt.Print("\n")
	//}

}
