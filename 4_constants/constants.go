package main

import "fmt"

// global scope me := ka use nhi kr skte hai
const age = 30 // ye chlega  but we can't do  age:=40 ..this is not allowed 
//baki hum normal variables declare kr skte hai := ye use krke nhi kr paa rhe hai 
// division := "first" 

var division = "first" //ye chal jaegi 

func main (){
	const name string = "golang" // const name = "golang"

	// name = "javascript" // yaha pr hum constant ki value ke sath chhed chhad nhi kr skte hai 

	fmt.Println(division)
	fmt.Println(name)
	fmt.Println(age)

	// when we are going to use constant for multiple variables then we use 
	const (
		port = 5000
		host = "localhost"
	)

	fmt.Println(port, host)
	

}
