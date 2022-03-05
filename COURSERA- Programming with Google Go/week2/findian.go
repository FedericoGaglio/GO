

package week2

import (
	"fmt"
	stringUtil "strings"
)

func FindIAN() {
	var inputUserString string

	fmt.Printf("Write you string:")

	_, err := fmt.Scan(&inputUserString)

	if err != nil {
		fmt.Printf("Input errato!")
	}

	validWord := inputUserString

	if( (stringUtil.HasPrefix(validWord, "i") || stringUtil.HasPrefix(validWord, "I")) &&
		(stringUtil.Contains(validWord, "a")  || stringUtil.Contains(validWord, "A")) &&
		(stringUtil.HasSuffix(validWord, "n")  || stringUtil.HasSuffix(validWord, "N"))){
			fmt.Printf("Found! \n")
	} else {
		fmt.Printf("Not Found! \n")
	}
	
}
