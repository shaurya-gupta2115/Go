package main

import "fmt"

func main(){

	var name string = "Shaurya Gupta"
	// jb tk name use nhi krenge to ya to delelte krna hoga nhi to usse use krna hoga file me 
	fmt.Println(name)

	// agar hum type nhi dete haito wo khud apna type select krleta hai during assingment of variable

	var isAdult = true // when hover kroge to pata chl jaega ki bool hai i nhi 

	fmt.Println(isAdult)

	var age int = 30
	fmt.Println(age)

	//shorthand syntax
	another_name := "goland" // ye shorthand use krte hai jb variable ko declare bhi krni hai and uski value bhi assing krni hoti hai to 
	fmt.Println(another_name)



	// agar kewal start me define bs krna hai aur baad me assingment krna hai to hume tabhi batana pdega ki kis type ka hai 
	var newName string
	newName = "Lucky Chauhan"

	fmt.Println(newName)


	// var price float32 <== first way
	// price = 50.3

	// var price = 50.3 <== second way

	price := 50.3
	fmt.Println(price)

	

}

