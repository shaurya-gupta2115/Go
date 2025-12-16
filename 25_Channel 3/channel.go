package main

import "fmt"


func task(done chan bool){

	defer func(){done <-true}()
	fmt.Println("Processing ....")
	// done <- true //agar error uske pehle ho gya to ye to chlega hi nhi 

	/*
		Agar panic / return / error aaya:
		❌ Signal nahi jayega
		❌ main forever block
		❌ DEADLOCK

		but jb defer use kiya hai to 

		✔ Function exit hone se pehle chalega
		✔ Panic ho ya return — signal jayega
		✔ Safe completion signal

		🔴 Channel = Communication + Signal
		🔵 WaitGroup = Sirf Signal

		jb bhi ek goroutines wagairah hote hai to hum use channel ke through krna chahte hai lekin agar
		kuchh extra implement karna hai to hum waitgroups ka use kr skte hai 

	*/


}


func main(){

	done := make(chan bool)
	go task(done)

	<- done //abhi ye block ho gya hai and exit ho jaega

}