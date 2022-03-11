package week1

import "fmt"

func BubbleSort(integers []int){
	boolArrived := true
	for boolArrived {
		boolArrived = false
		for i:=0 ; i<len(integers)-1 ; i++ {
			if(integers[i] > integers[i+1]){
				Swap(integers,i)
				boolArrived = true
			}
		}
	}
	fmt.Print(integers)
}

func Swap(integers []int, index int) {
	temp := integers[index]
	integers[index] = integers[index+1]
	integers[index+1] = temp
}
