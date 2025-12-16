package main

import "fmt"

// func printSlice(items [] int){
// 	for _, item := range items{
// 		fmt.Println(item)
// 	}
// }

// func printStringSlice(items []string){
// 	for _, item := range items{
// 		fmt.Println(item)
// 	}
// }


//to avoid this duplicacy 
func printSlice[T any](items [] T){ //[T any] using any is not a good practise any == interface{} 
	for _, item := range items{
		fmt.Println(item)
	}
}

func anotherPrintSlice[T int | string](items [] T){
	for _, item := range items{
		fmt.Println(item)
	}
}

type stack[T int|string] struct {
	elements []T
}

func main(){

	myStackint := stack[int]{
		elements: []int{2,5,25,2,1},
	}
	fmt.Println(myStackint) // ye string ka struct de dega

	myStackString := stack[string]{
		elements: []string{"hello", "world"},
	}
	fmt.Println(myStackString)

	nums := []int{4,6,3,6,27,2,1,46,346,3}
	names := []string{"golang", "typescript"}
	// printSlice(nums)
	// printStringSlice(names) 

	// Reverse the names slice manually
	// for i := len(names) - 1; i >= 0; i-- {
	// 	fmt.Println(names[i])
	// }

	printSlice(nums) // ab ye generic  bana gya hai dono ke liye to code repitition km ho gyi hai
	// ab yaha 'names' bhi dalenge to bhi koi dikkat nhi hogi 


	anotherPrintSlice(names) // now this will only accept int & string only




}