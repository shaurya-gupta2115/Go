package main

import "fmt"

func main(){
	//arrays sequence of specific length

	// zeroed value type ke hisaab se hi ho jati hai print

	var nums[4] int 
	nums[0] =1

	fmt.Println(nums[0])
	fmt.Println(nums) // by defualt sabhi ko 0 le leta hai -> depends on which type of array it is

	var vals[4] bool
	fmt.Println(vals) //by default false le liye hia 

	var st[3] string
	st[0] = "golang"
	fmt.Println(st) // empty array string ans sbme "" ye hi hoga

	fmt.Println(len(nums))

	// agar declaration bhi sath sath me hi dena hai to  => declaring in single line
	ram := [3]int {1,2,3}
	fmt.Println(ram)

	// 2D arrays
	array := [2][2] int {{3,4}, {5,6}}
	fmt.Println(array) //  => fixedsize , memmory optimzation, time is predicatable 

	// instead of using arrays we generally use slices , or we can say slices are the vectors of the c++
	// for dynamically size extension :)

}
