package main

import "fmt"

func main() {
	// iterating over data structures
	nums := []int{6, 7, 8}

	// for i:= 0; i<len(nums);i++{
	// 	fmt.Println(nums[i])
	// }

	// sum := 0
	for index, num := range nums { // it like auto i : nums to i me saari nums ki value jati hai similarly hai ye bhi
		// sum = sum + num
		fmt.Println(num, index)
	}
	// fmt.Println(sum)

	//lets take a map
	m := map[string]string{"fname": "Shaurya", "lname": "Gupta"}

	for k, v := range m {
		fmt.Println(k, v) // fmt.Println(i+"-"+str) is trh bhi kr skte hai hum
	}

	for k := range m {
		fmt.Println(k) // agar kewal ek hi argument dete hai to bydefault wo key ko hi lega
	}

	for i, c := range "golang"{
		// fmt.Println(i,c)// index unicode(ASCII VALUE)
		fmt.Println(i,string(c))// index unicode(ASCII VALUE) converted to character using string 
	}


}
