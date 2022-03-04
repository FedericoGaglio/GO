

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


//if(!(validWord[0]=='i' || validWord[0]=='I') ||
	   //!(validWord[len(validWord)-1] == 'n' || validWord[len(validWord)-1] == 'N') ) {
		//fmt.Printf("Not Found! \n")
	//}

	//if((validWord[0]=='i' || validWord[0]=='I') &&
	   //(validWord[len(validWord)-1] == 'n' || validWord[len(validWord)-1] == 'N') ) {

		
		// il range ritorna due valori:
	//PRIMO: la posizione dell'attuale valore che si sta scorrendo
	//SECONDO: una copia del valore effettivo


		//for pos, _ := range validWord[1:len(validWord)-2] {
			//if(validWord[pos] == 'a' || validWord[pos] == 'A'){
				//fmt.Printf("Found! \n")
				//break
			//}
		//}
	//}