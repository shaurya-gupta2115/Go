package main

import "fmt"

// go me closure bahut powerfull hote hai. wo ek aisa function hota hai jo apne surrounding (outer) function ee variables ko "yaad"
//rakh skta hai , chahe outer function ka execution khtam ho chuka ho .

//Ye concept bahut powerful hai, especially jab aapko kisi function ke andar state yaad rakhni ho.
func counter() func() int{
	count := 0
	return func() int{
		count+=1
		return count
	}
}

func main (){
	increment := counter() // increment bhi ek function hai 

	fmt.Println(increment())
	fmt.Println(increment())
	fmt.Println(increment())

}

// Important Points

// Closure inner function hai jo outer function ke variables ko access kar sakta hai.
// Har bar jab closure ko call karte ho, wo apni state maintain karta hai.
// Closures ko aap stateful function bana sakte ho.


// Summary
// Closure = function + uska environment (variables)
// Useful for: state yaad rakhna, data hiding, functional programming patterns