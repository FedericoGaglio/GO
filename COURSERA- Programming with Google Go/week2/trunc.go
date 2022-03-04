

package week2

import (
	"fmt"
)

func Trunc() {
	var inputUserVariable float64

	fmt.Printf("Write the number please:")

	_, err := fmt.Scan(&inputUserVariable)

	if( err!= nil) {
		fmt.Printf("Errore nella digitazione del valore da input \n")
	}

	intValue := int64(inputUserVariable)

	fmt.Printf("The converted number is %v \n", intValue)
}