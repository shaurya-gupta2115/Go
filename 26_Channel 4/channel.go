package main

import (
	"fmt"
	"time"
)

//non-blocking / buffer based channel

func emailSender(emailChan chan string, done chan bool){
	defer func(){done <- true}()

	for email := range emailChan{
		fmt.Println("Sending email to : " , email)
	}
}

func main(){
	emailChan := make(chan string, 100) // think like 100 items desk me rakh skte hai 
	done := make(chan bool) //ye humne banaya hai syncronise krne ke liye 

	go emailSender(emailChan, done)

	for i:=0; i<=99; i++{ // 100 capacity thi to uske baad block krne dega
		emailChan <- fmt.Sprintf("%dgupta@gmail.com", i)
		time.Sleep(time.Millisecond *100)
	}

    close(emailChan) // very important: receiver ke range loop ko khatam karega

	<-done //wait till emailSender finishes

	// emailChan <- "shaurya@gmail.com"
	// emailChan <- "aditya@gmail.com"

	// fmt.Printf(<-emailChan)
	// fmt.Printf(<-emailChan)

}