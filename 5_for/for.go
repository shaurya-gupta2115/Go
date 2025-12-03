package main

import "fmt"


// for -> only construct in go for looping
//and isi ki implementation ka use krke hi hume while loop banana / simulate krna hota hai 

func main(){

	//while loop
	i:=1 
	for i<=3 {
		fmt.Println(i)
		i++
	}

	fmt.Println("While loop ended and for loop started")
	// classical for loop 
	for j:=0; j<=3; j++{

		if j ==2 {
			continue
		}
		fmt.Println(j)
	}

	fmt.Println("Range - started")

	// range function newly introduced

	for k:= range 11{
		fmt.Println(k)
	}

}
