package main

import (
	"fmt"
	"time"
)


func task(id int){
	fmt.Println("Doing task :", id)
}

func main (){
	// Goroutine = Lightweight function that runs concurrently, managed by Go runtime

	for i:=0; i<=10; i++{
		// go task(i)

		go func(){
			fmt.Println(i)
		}()
	} 
	fmt.Println("after go routines")
	time.Sleep(time.Second * 2)


}

/*
Chef (Main Goroutine)
 |
 |-- Chop vegetables (goroutine)
 |-- Boil rice (goroutine)
 |-- Fry paneer (goroutine)

Chef sab ka wait nahi karta
Kaam parallel chalte hain


syntax : go functionName()

*/