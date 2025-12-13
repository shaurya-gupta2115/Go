package main

import "fmt"

func add(a int, b int) int {
	return (a + b)
}

//	func getlanguages()(string, string, string){
//		return "golang", "javascript", "c"
//	}
func getlanguages() (string, string, bool) {
	return "golang", "javascript", true
}

func processIt(fn func(a int) int) { // abhi humne ek function pass kara 'fn' jo 'a' int type ka lega and return krta hoga ek 'int'
	fmt.Println(fn(10))
}

func fn(a int) int {
	return a
}

func main() {

	// result := add(2,5)
	// fmt.Println(result)

	// fmt.Println(getlanguages()) //golang javascript c

	// lang1, lang2, lang3 := getlanguages()
	// fmt.Println(lang1)
	// fmt.Println(lang2)
	// fmt.Println(lang3)

	lang1, lang2, _ := getlanguages() //agr hume use hi do values krni hai aur nhi chahte ki 3rd value bhi koi store karaye to

	fmt.Println(lang1)
	fmt.Println(lang2)

	// Go me function ko ek variable me assign kar sakte hain:
	// Example:
	//   myFunc := add
	//   fmt.Println(myFunc(2, 3)) // Output: 5
	// Yaha 'add' function ko 'myFunc' variable me assign kiya aur usko call kiya.

	// Kisi bhi function ko as an argument dusre function me pass kar sakte hain:
	// Example:
	//   func processIt(fn func(int) int) {
	//       fmt.Println(fn(10))
	//   }
	//   processIt(fn) // Yaha 'fn' function ko argument ke roop me diya

	// Function se naya function bhi return kara sakte hain:
	// Example:
	//   func makeAdder(x int) func(int) int {
	//       return func(y int) int {
	//           return x + y
	//       }
	//   }
	//   adder := makeAdder(5)
	//   fmt.Println(adder(3)) // Output: 8
	// Yaha 'makeAdder' function ek naya function return karta hai jo 'x' ko yaad rakhta hai (closure)
	//Jab aap function banate ho jo ek function return karega, toh uska return type bhi function ka signature hota hai.
	//Yaha makeAdder ek function hai jo ek int leta hai aur return karta hai ek function jo int leta hai aur int return karta hai.

	processIt(fn)
}
