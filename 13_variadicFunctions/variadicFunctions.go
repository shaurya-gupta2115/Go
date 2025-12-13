package main

import "fmt"

//Variadic function wo hota hai jisme aap ek hi parameter me multiple (0 ya zyada) values pass kar sakte ho.

// nums ...int ka matlab hai ki function me nums ek slice ban jayega ([]int).
// Agar aapke paas already ek slice hai, toh usko variadic function me pass karne ke liye ... lagana padta hai:

func sum(nums ...int)int{ // aur jb multiple type ke datatypes hote hai to interface ka use krte hai sum(nums ...interface{})
	total :=0
	for _,n := range nums{
		total += n;
	} 
	return total
}
func printAll(prefix string, nums ...int) { // variadic function 
    fmt.Print(prefix)
    for _, n := range nums {
		fmt.Println(n)
    }
    fmt.Println()
}

func main(){ 
	
	addition := sum(4,5,63,637,421,64,44)
	fmt.Println(addition)
	fmt.Println()

	// what if there is array /list
	nums := []int{1,2,3,4,5}
	result := sum(nums...) // ye ... lagana jaruri hai agr array/list pass kar rahe ho to
	// result := sum(nums)//agar ... nhi lagayenge to error aayega kyunki sum function me int type ke multiple values chahiye na ki ek slice
	fmt.Println(result)

	// we can also pass multiple types of data using interface{}
	fmt.Println(1,2,3,4,5,"hello")
	printAll("Numbers :", 10,20,30,40,50)
}
