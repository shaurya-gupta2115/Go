package main

import "fmt"

func main(){

	fmt.Println("hey! I am starting!")
	for i:= 0; i<10 ;i++{
		defer fmt.Println(i)
	}
	fmt.Println("done")
	// concept hai yaha pr stack ka pehle sb priint ho jaenge and through out the process wo stack me padta jaega
	// lekin uske baad se jb function exit hoga tb print hoga
	// to yaha pr last me 9 print hoga sbse phle and first me 0 print hoga sbse last me : LIFO ORDER me 


}   