package week4

import (
	"fmt"
	"encoding/json"
	"log"
)

func MakeJSON() {

	var userInputName string
	var userInputAddr string

	var infoMap map[string] string
	infoMap = make(map[string] string)

	fmt.Print("Please digit a name : \n")
	_, err := fmt.Scan(&userInputName)

	if(err != nil){
		fmt.Print("Error whit input ! \n")
	}

	fmt.Print("Please digit an address : \n")
	_, err2 := fmt.Scan(&userInputAddr)

	if(err2 != nil){
		fmt.Print("Error whit input ! \n")
	}

	infoMap[userInputName] = userInputAddr

	mapDataToJSONFormat, err := json.Marshal(infoMap)

	if err != nil {
		log.Println(err)
	}
	fmt.Println(string(mapDataToJSONFormat))


}