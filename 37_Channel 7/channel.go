package main

func sum(nums []int, c chan int){
	total := 0
	for _, v := range nums {
		total += v
	}
	c <- total
}

func main(){
	nums := []int{7,2,8,-9,4,0}
	c := make(chan int)

	// hum kisi work ko concurrently do alag alag goroutines me kr skte hai

	// jb hum goroutine banate hai to wo function ko call krta hai aur uske andar ka code concurrently execute hota hai
	// iska mtlb ye hai ki ye code main thread se alag thread me chalta hai aur dono ek sath chalte hai

	// jb hum goroutine banate hai to hume uske liye koi alag se thread create krne ki jarurat nhi pdti hai go runtime khud hi manage krta hai
	// jb hum goroutine banate hai to wo immediately execute hona start ho jata hai

	//jb order matter nhi krta hai to hum goroutine ke work ko divide kr skte hai 
	go sum(nums[:len(nums)/2], c)
	go sum(nums[len(nums)/2:], c)

	x,y := <-c, <-c // receiving from channel is a blocking operation

	// jb tak dono values channel se receive nhi ho jati tab tak ye line aage nhi badhegi
	// iska mtlb ye hai ki jb tak dono goroutines apna kaam complete karke channel me value send nhi kr deti tab tak ye line wait krti rahegi
	
	println(x, y, x+y)
}