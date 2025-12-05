package main

import (
	"fmt"
	"slices"
)

// now we are going to deal with slices => dynamic in nature and highly used throughout the go
// we do not give the size => most used construct + useful methods

func main(){

	//uninitialised slice is nil (null)
	var nums[] int
	fmt.Println(nums == nil) //true 
	fmt.Println(len(nums))
	fmt.Println(nums)
	// true
	// 0
	// []

	var num = make([]int, 3,5) // datatype , initial size, initial capacity
	fmt.Println(len(num)) // 3
	fmt.Println(num) // [0,0,0]
	fmt.Println("Capacity is ", cap(num)) // capacity is 5
	
	num = append(num, 3)
	num = append(num, 5)
	num = append(num, 6)
	fmt.Println(num) // [0 0 0 3 5 6]
	fmt.Println("Capacity is ", cap(num)) // Capacity is  10

	num = append(num, 398)
	num = append(num, 534)
	num = append(num, 894)
	num = append(num, 574)
	num = append(num, 688)
	fmt.Println("num =" , num) // [0 0 0 3 5 6 398 534 894 574 688]
	fmt.Println("Capacity is ", cap(num)) // yaha pr sb double hota jaa rhi hai capacity jb size se jyada hum elements dal de rhe hai 

	var sli = make([]int, 0, 5) //yaha pr maine initialise kiya hai and then by default size 0 haiii but capacity abhi hai 5 ki 
	// to isse kya hota hai agar hum size de dete certain value to wo nil nhi kehlati but agar hum abhi compare kare to abhi nil dikhaega

	//copy function 
	copy(num, sli)
	fmt.Println("slice sli = ", sli)

	//:::: now to avoid confusion in array and slice  ::::
	var a[3] int //zero value of array of length 3 => [0,0,0]
	b := [3] int{1,2,3} //explicit array of length 3
	c := [...]int {2,35,64,23,61,44} // yaha pr wo length apne aap bhar leta hai 
	fmt.Println(a, "   ", b, "   " , c)

	// for slice representation we go like this 

	var d[] int // ye declare krdega but abhi ye nil hi rhega ...it is not like []
	if d == nil{ //abhi ye true dega because abhi tautological condition hai ye 
		fmt.Println(true)
	}
	s := []int {3,46,63,2,3,665,35,23}

	// different methods to create slices 
	// var s []int              // nil slice, length 0, capacity 0
	t := []int{1, 2, 3}      // slice literal (backed by an array internally)
	u := make([]int, 5)       // slice with len=5, cap=5 (filled with zero values)
	v := make([]int, 0, 10)   // len=0, cap=10 (reserved capacity)
	w := append(s, 7)         // append works on slices

	fmt.Println(d , "     " , s, " ", t, " " , u , " " , v, " ", w)

	// comparing two slices
	sl1 := []int{1,2,43}
	sl2 := []int{1,2,43}
	fmt.Println(slices.Equal(sl1, sl2))

	//2d slice creation 
	var newnum = [][] int{{1,2,3}, {4,5,6}}
	fmt.Println(newnum)


	//comparison and assignment of the array and slices 
	first := [3]int{1,2,3}
	second := first
	second[0] = 9
	fmt.Println(first) // still [1 2 3] -> arrays copied

	third := []int{1,2,3}
	fourth := third
	fourth[0] = 9
	fmt.Println(third) // [9 2 3] -> slices share underlying array
}


// Why Go designers separate them? (intuition)
// 	•	Arrays: fixed-size, efficient when you always know exact size (rare in real apps). Useful for low-level, small fixed buffers.
// 	•	Slices: flexible, idiomatic for most programs. They are backed by arrays but provide dynamic behavior (append, reslice, etc.).

// make([]T, n) or make([]T, n, cap) → slice creation (dynamic).
// •	[]T{...} → slice literal (but note: []T{} is slice; [N]T{} or [...]T{} is array).
// •	arr[:] or arr[1:3] → slicing an array returns a slice.
// •	for i := range s { ... } where s is []T → slice iteration.
// •	var arr [1000]byte for fixed buffer -> array.
// •	buf := make([]byte, 0, 1024) -> slice preallocated for efficiency.


// a := [3]int{1,2,3}
// b := a
// b[0] = 9
// fmt.Println(a) // still [1 2 3] -> arrays copied

// s := []int{1,2,3}
// t := s
// t[0] = 9
// fmt.Println(s) // [9 2 3] -> slices share underlying array