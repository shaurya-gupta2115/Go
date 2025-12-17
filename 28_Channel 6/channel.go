package main

import (
	"fmt"
	"time"
)

func emailSender(emailChan <-chan string, done chan<- bool){
	defer func(){ done <- true }()

	// emailChan <- "hello" // hum ye nhi kr skte hai kyunki hum channel me send kr rhe hai and since it is receive only channel hai to hum send krna allow hi nhi hai 

	// <-chan  => it means that the channel ke upr receive bs kr skte hai  
	// chan<-  => this means that the channel ke upr send bs kr skte hai, receive nhi kr skte 

	for email := range emailChan{
		fmt.Println("Sending email to ", email)
		time.Sleep(time.Second)
	}
}

func main() {

	chan1 := make(chan int)
	chan2 := make(chan string)

	go func() {
		chan1 <- 10
	}()

	go func() {
		chan2 <- "hello, goroutines"
	}()

	for i := 0; i < 2; i++ {
		select {
		case chan1Val := <-chan1:
			fmt.Println("received data from chan1 :", chan1Val)

		case chan2Val := <-chan2:
			fmt.Println("received data from chan2 :", chan2Val)
		}
	}

}
