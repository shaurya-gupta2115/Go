package main

import (
	"fmt"
	"sync"
	"time"
)

func main(){
	var wg sync.WaitGroup
	
	go func(){
		defer wg.Done()
		fmt.Println("Work Started")
		time.Sleep(time.Second *2)
		fmt.Println("Work Done")
	}()

	wg.Wait() // yaha tb tak waiting state me rahege jb tk sare goroutines done na karde 
}
