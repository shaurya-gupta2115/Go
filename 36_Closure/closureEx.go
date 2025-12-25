package main

/*
LINE-BY-LINE EXECUTION EXPLANATION (Closure)

Code:

	func add(x int) func(int) int {
		sum := 0
		return func(y int) int {
			sum += (x + y)
			return sum
		}
	}

1) Function signature ka meaning

	func add(x int) func(int) int

	- add ek input leta hai: x (type int)
	- add return karta hai ek function type: func(int) int
	  yani “ek function jo ek int input lega (y) aur int return karega”.

2) main() me pehli line:

	add5 := add(5)

	Step-by-step:
	- add(5) call hota hai => x = 5
	- sum := 0 create hota hai (initially 0)
	- add ek anonymous function return karta hai:

		  func(y int) int {
			  sum += (x + y)
			  return sum
		  }

	  Ye returned function “closure” hai, kyunki ye outer variables x aur sum ko capture/remember karta hai.
	  add() ka execution end ho jaata hai, but x aur sum alive rehte hain because add5 unko hold karta hai.

3) Ab add5(...) calls

	(A) println(add5(3))
		- y = 3
		- current: x=5, sum=0
		- (x+y) = 5+3 = 8
		- sum = sum + 8 = 0 + 8 = 8
		- return 8 => print 8

	(B) println(add5(7))
		- y = 7
		- IMPORTANT: sum reset nahi hota; same sum carry forward hota hai
		- current: x=5, sum=8
		- (x+y) = 5+7 = 12
		- sum = 8 + 12 = 20
		- return 20 => print 20

	(C) println(add5(10))
		- y = 10
		- current: x=5, sum=20
		- (x+y) = 5+10 = 15
		- sum = 20 + 15 = 35
		- return 35 => print 35

So, actual output (is code ke hisaab se):

	8
	20
	35

NOTE:
File me jo inline comments “// 15” aur “// 25” likhe hain, wo current logic ke saath match nahi karte.
Current logic har call me (x+y) ko sum me add karta hai, isliye values 8, 20, 35 aati hain.
*/

func add(x int) func(int) int {
	sum := 0
	return func(y int) int {
		sum += (x + y)
		return sum
	}
}

func main() {
	add5 := add(5)
	println(add5(3))  // 8
	println(add5(7))  // 15
	println(add5(10)) // 25
}
