package main

import "fmt"

//pass by value 
func changeNumValue(num int) {
	num = 5
	fmt.Println("value in function : ", num)

}

// pass by reference 
func changeNumRef(num *int){
	*num = 5
	fmt.Println("value in function : ", *num)
}

func main() {
	num := 1
	changeNumValue(num) 
	fmt.Println("After change num: ", num) // abhi ye pass by value hua hai to dekha hai yaha pr koi actual wale me changes nhi hue hai 

	newnum :=11
	changeNumRef( &newnum)
	fmt.Println("After change newnum: ", newnum) // abhi ye pass by value hua hai to dekha hai yaha pr koi actual wale me changes nhi hue hai 
}
