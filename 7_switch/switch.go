package main

import (
	"fmt"
	"time"
)

func main(){
	//simple switch

	i:= 3

	switch i {
	case 1: fmt.Println("Value is one")
	case 2: fmt.Println("Value is two")
	case 3: fmt.Println("Value is three")
	default: fmt.Println("other")
	}

	///multiple condition swithc 
	switch time.Now().Weekday(){
	case time.Saturday, time.Sunday :
		fmt.Println("it's weekend")
	default: 
		fmt.Println("it is workday")
	}

	// type switch 
	// interface{} means datatype is 'anytype' koi bhi ho skta hai
	whoAmI := func(i interface{}){
		switch t := i.(type){			//switch := i.(type)  when we don't want to use other variable assingment 
		case int: 
			fmt.Println("given input is integer")
		case float32: 
			fmt.Println("given input is float")
		case string: 
			fmt.Println("given input is string")
		default: 
			fmt.Println("other", t)
		}
	}

	whoAmI(5)
}
