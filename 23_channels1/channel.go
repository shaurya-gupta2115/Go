package main

import (
	"fmt"
	"math/rand"
	"time"
)

//Channel is a typed pipe used to send and receive data between goroutines

func processNum(numChan chan int) {
	for num := range numChan {
		fmt.Println("Processing number: ", num)
		time.Sleep(time.Second)
	}
}

func main() {
	/*
		Unbuffered channel me send/receive dono ko same waqt par ready hona padta hai (synchronous handshake).
		Tumhare code me send ke time koi receiver ready nahi, isliye deadlock.
	*/

	// messageChan := make(chan string,1) //yaha pr buffer size >=1 hona chahiye tabhi ek box ki jgh ek self ki trh work krega and deadloack nhi aaega
	// messageChan <- "ping" // (sending) into channel & blocking rhega jb tk receive nhi ho jata hai tb tk
	// msg := <- messageChan // (receiver)
	// fmt.Println(msg)

	numChan := make(chan int)
	go processNum(numChan) // ye non blocking hai abhi kyunki 2nd goroutine bana hua hai

	// numChan <- 5 // abhi humne channel me 5 dala hai lekin ye turnt run kr jaega to processNum to kuchh print hi nhi kr paega

	for {
		numChan <- rand.Intn(100)
	}

	// time.Sleep(time.Second *2 )

}

/*
	Kya main bhi goroutine hota hai?

	Jab go run ya go build se program start hota hai, Go runtime ek default goroutine start karta hai jo main.main() function chalata hai.
	Ye bhi ek normal goroutine hi hai, bas iska naam hum usually “main goroutine” bol dete hain.
	Jab main() return ho jaata hai, poora program khatam ho jaata hai, chahe baaki goroutines chal rahe ho ya nahi.

	func main(){
		messageChan := make(chan string)

		messageChan <- "ping"
		msg := <- messageChan

		fmt.Println(msg)
	}

	Sirf ek hi employee hai = main goroutine.
	Wo khud hi dabbe me "ping" daalke kisi ko dena chahta hai:
	messageChan <- "ping"
	Unbuffered rule: jab tak saamne koi receiver ready nahi, ye employee ruka rahega (block).
	Problem:
	Yehi employee ko aage jaake receive bhi karna hai (msg := <-messageChan), lekin wo send pe hi atak gaya.
	Koi second employee (dusra goroutine) hai hi nahi jo receive kare.

	Result:
	Saare employees (yaha 1 hi hai) ruk gaye / so gaye → runtime bolta hai:
	all goroutines are asleep - deadlock.

	Deadlock exactly kya hai yaha?

	Deadlock = aisi situation jahan koi bhi goroutine aage nahi badh sakta, sab kuch kisi na kisi resource/channel pe block hai.
	Tumhare case me:
	Main goroutine channel pe send pe block.
	Koi goroutine receive karne ko exist hi nahi karta.
	Runtime check karta hai: “koi runnable goroutine bacha?”
	Answer: nahi → deadlock error.

	func main() {
		messageChan := make(chan string)
		go func() {
			// yeh second goroutine hai
			messageChan <- "ping"
		}()
		// yaha main goroutine receive karega
		msg := <-messageChan
		fmt.Println(msg)
	}

	Analogy:
	Ab office me 2 employees:
	Employee A (main goroutine) receive karega.
	Employee B (new goroutine) send karega.
	Jab B dabba dega, A ready hai lene ke liye → koi block nahi, koi deadlock nahi.

	func main() {
    messageChan := make(chan string, 1) // buffer size 1
    messageChan <- "ping"  // yeh ab block nahi hoga
    msg := <-messageChan
    fmt.Println(msg)
	}

	Analogy:

	Channel ab ek dabba nahi, ek shelf hai jisme 1 item rakh sakte ho.
	Sender ne "ping" shelf pe rakh diya, usko turant receiver nahi chahiye.
	Baad me same employee (main goroutine) jaa ke shelf se uthata hai.
	Kyunki shelf me jagah hai, send time pe koi handshake zaroori nahi → no deadlock.


*/
