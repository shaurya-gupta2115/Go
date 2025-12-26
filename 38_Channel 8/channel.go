package main


func fibonacciCap(n int, c chan int){
	x, y := 0, 1
	for i := 0; i<n ;i++{
		c<- x
		x, y = y, x+y
	}
	close(c) // channel ko close krna jaruri hai jb hum buffered channel use kr rhe ho
}

func main(){
	c := make(chan int, 10) // buffered channel

	// ab hum is channel me 10 values bhej skte hai bina kisi blocking ke
	// jb tak channel full nhi ho jata tab tak hum isme values bhejte rah skte hai

	go fibonacciCap(cap(c), c)

	for v:= range c{
		println(v)
	}

	// bs do cheez yaad rakhni hai
	// 1. buffered channel me values bhejna tab tak block nhi hota jab tak channel full na ho jaye
	// 2. buffered channel se values receive karna tab tak block nhi hota jab tak channel empty na ho jaye

	// and hume close ka dhyan rakhna hai taaki jb hum buffered channel use kr rhe ho to usse close krna na bhool jaye
	//close helps to inform main goroutine that no more values will be sent on this channel
	
	//range method se hum channel se values receive kr skte hai jb tak channel close na ho jaye
}