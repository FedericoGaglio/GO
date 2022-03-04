
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

	// creazione di uno slice di interi di dimensione 3
	elementsSlice := make([] int, 3)

	inputScanner := bufio.NewScanner(os.Stdin)

	fmt.Printf("Hi user, ready to digit an integer number =) ? \n")

	fmt.Printf("Please digit a number: \n")

	//QUANDO USI LA FUNZIONE SLICE IN UN CICLO FOR RICORDA CHE:
	// il range ritorna due valori:
	//PRIMO: la posizione dell'attuale valore che si sta scorrendo
	//SECONDO: una copia del valore effettivo
	for inputScanner.Scan() {

		fmt.Printf("Please digit a number: \n")

		userInputStr = inputScanner.Text()
		userInputInt, err := strconv.Atoi(userInputStr) // converting string to int

		if strings.EqualFold(userInputStr, "X") { // se metto il carattere di uscita allora brekko ed esco

			fmt.Printf("Ok, you decide to terminate de number insertion, Bye \n")
			fmt.Printf("This is your sorted slice \n")
			break
		}

		if (err != nil) { // se metto una stringa come inpuit allora brekko ed esco

			fmt.Printf("Err on given input! \n")
			fmt.Printf("This is your slice of integers \n")
			break
		}

		elementsSlice = append(elementsSlice, userInputInt)

	}

	elementsSlice = OrderSlicer(elementsSlice)
	elementsSlice = elementsSlice[3:] //cancello le prime te passa che inseriscono di fatto 3 zeri

	fmt.Print(elementsSlice)
}

func OrderSlicer( passedSlice[] int ) [] int {
	

	//dichiaro la mappa che di fatton mi permettera di poter andare a fare l'ordinamento dello slice passato
	var mapAppo map[int] int
	mapAppo = make(map[int] int)

	//for key,value := range passedSlice {
		//mapAppo[pos] = passedSlice[pos]
	//}

	//sulla base deglj elementi dello slice pescati per posizione , gli inserisco mettendo quindi nella mappa un inserimento
	//chiave-valore (ex: posizione index 1 diventa chiave "1", con value slice[1], etc...)
	for actualPos,_ := range passedSlice {
		mapAppo[actualPos] = passedSlice[actualPos]
	}

	//creo uno slice di interi di dimensione identica alla mappa creata in precedenza
	valuesOfTheMap := make([] int, 0, len(mapAppo))

	//scorro la mappa andando a prendere in considerazione unicamente i valori di tale mappa e andando a "mutare" nel for le keys
	//inserisco poi questi VALUES nel mio nuovo slice di interi appena creato
	for _,values := range mapAppo {
		valuesOfTheMap = append(valuesOfTheMap, values)
	}

	//ordino grazie alla libreria sort i valori interi
	sort.Ints(valuesOfTheMap)

	return valuesOfTheMap
}