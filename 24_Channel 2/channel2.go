package main

import "fmt"

// doing channel 2 example and its implementation

func sum(result chan int, num1 int, num2 int){
	newResult := num1 + num2
	result <- newResult
}

func main(){ //OS → Go runtime → main goroutine
	// creating new channel  
	result := make(chan int)

	// starting the goroutine
	go sum(result, 4, 6) // agar go nhi likhe to parallely execute nhi hoga jiske krn newResult tb hi daal payega jb koi recieve krne ke liye ready ho ya fir handshake krne ke liye available ho isliye parallely execute karana important step hai taki res <- result ko le paye 


	res := <-result //result channel ka data res me enter ho paye


	fmt.Println(res)

}


/*
Runtime : 
	result channel
	type: chan int
	buffer: 0 (unbuffered)

	go sum(result, 4, 6)


	main goroutine
	|
    |──▶ new goroutine created
            |
            v
        sum()

	G1 → main
	G2 → sum

	newResult := num1 + num2

	result <- newResult

	G2 (sum)  ---> [ CHANNEL ] <---  G1 (main)
    send                               receive

	result <- newResult  // completed
	return
	G2 → EXIT
	res := <-result   // res = 10
	fmt.Println(res)


	Time →
	main:   go sum ──recv(wait)────────print
	sum:          add ─send(wait)─exit

	Go me goroutine: =>directly return nahi kar sakta => channel ke through return karta hai



*/