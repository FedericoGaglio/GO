
package week3

import (
	"fmt"
	"sort"
	"strconv"
	"bufio"
	"os"
	"strings"
)

func SliceInsert() {
	var userInputStr string

	elementsSlice := make([] int, 3)

	inputScanner := bufio.NewScanner(os.Stdin)

	fmt.Printf("Hi user, ready to digit an integer number =) ? \n")

	fmt.Printf("Please digit a number: \n")

	for inputScanner.Scan() {

		fmt.Printf("Please digit a number: \n")

		userInputStr = inputScanner.Text()
		userInputInt, err := strconv.Atoi(userInputStr) 

		if strings.EqualFold(userInputStr, "X") { 

			fmt.Printf("Ok, you decide to terminate de number insertion, Bye \n")
			fmt.Printf("This is your sorted slice \n")
			break
		}

		if (err != nil) { 

			fmt.Printf("Err on given input! \n")
			fmt.Printf("This is your slice of integers \n")
			break
		}

		elementsSlice = append(elementsSlice, userInputInt)

	}

	elementsSlice = OrderSlicer(elementsSlice)
	elementsSlice = elementsSlice[3:] 

	fmt.Print(elementsSlice)
}

func OrderSlicer( passedSlice[] int ) [] int {
	
	var mapAppo map[int] int
	mapAppo = make(map[int] int)

	for actualPos,_ := range passedSlice {
		mapAppo[actualPos] = passedSlice[actualPos]
	}

	valuesOfTheMap := make([] int, 0, len(mapAppo))

	for _,values := range mapAppo {
		valuesOfTheMap = append(valuesOfTheMap, values)
	}

	sort.Ints(valuesOfTheMap)

	return valuesOfTheMap
}
