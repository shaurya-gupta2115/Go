package main

import "fmt"

func main(){

	age :=19
	if age>=18{
		fmt.Println("He / She is an adult")
	} else {
		fmt.Println("he is not adult")
	}

	//there is another way to initialise the if statement
	if name:="Shau"; len(name) >= 5{ 			// yaha same time pr initialise kiya hai name ko
		fmt.Println("yes the name has greater size")
	} else if len(name) <=2{
		fmt.Println("The size is extremely less")
	} else {
		fmt.Println("Size is small but can be managed ")
	}

	// go does not have ternary operator => we have to use normal if-else

}
