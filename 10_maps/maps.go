package main

import (
	"fmt"
	"maps"
)

func main(){
	// maps are key value pairs => like dictionary in python or objects in js	
	// make(map[KeyDatatype]ValueDatatype)

	m := make(map[string] string)
	//setting key value pairs
	m["name" ] = "golang" 
	m["version"] = "1.19"
	m["creator"] = "google"
	m["use"] = " backend"

	fmt.Println(m["name"])
	fmt.Println(m["creator"])

	fmt.Println(m["hello"]) // if key does not exist in the map , then it return nothing => zero value  
	//for stirng => blank line // for int => false


	// lets create another map
	fmt.Println("new map started")
	new := make(map[string] int)
	// new["hello"] = "kya hal hai " // this will throw error 

	new["aditya sharma"] = 22
	new["shaurya gupta"] = 22

	fmt.Println(new["aditya sharma"])
	fmt.Println(new["shaurya gupta"])
	fmt.Println(new["lucky chauhan"]) // this will print 0 , zero value for integer

	// to find the length
	// fmt.Println(len(new))

	delete(new, "shaurya gupta")
	fmt.Println(new["shaurya gupta"])

	clear(new)

	fmt.Println(new)

	// another way to implement the map function 
	mp := map[string] int{"price":45, "phones":3}
	fmt.Println(mp)

	//to check whether element i spresent or not in map 

	_, ok := mp["price"] // agar ok hoga to true aega , agar ok nhi hoga to false aaega

	if ok{
		fmt.Println("all ok") // agar _ ki jgh k hota hai to k =45 rhega price ki value nikal kr aati hai 
	}else{
		fmt.Println("not ok")
	}


	k, ok := mp["prices"] // agar ok hoga to true aega , agar ok nhi hoga to false aaega
	if ok{
		fmt.Println("all ok", k) // agar _ ki jgh k hota hai to k =45 rhega price ki value nikal kr aati hai 
	}else{
		fmt.Println("not ok", k)
	}

	m1 := map[string] int{"price":45, "phones":3}
	m2 := map[string] int{"price":45, "phones":3}
	
	// fmt.Println(m1== m2) this don't work in go since it is object
	fmt.Println(maps.Equal(m1, m2)) // object direct match nhi hota hai ...uske liye first 
	// maps me Equal function ka use krna hota hai 
	

}