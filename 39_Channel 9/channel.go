package main

/*
ENHANCED / DEEP EXPLANATION (Channels + Goroutines + select)

High-level picture:

There are 2 goroutines running:
	1) main goroutine: runs main(), then calls fibonacci(c, quit)
	2) anonymous goroutine: reads 10 values from channel c, prints them, then signals quit

There are 2 channels (both are UNBUFFERED):
	- c:   used for sending fibonacci numbers
	- quit: used for stop/cancel signal

UNBUFFERED channel ka key rule (handshake model):
	- Send (c <- x) tab tak BLOCK rahega jab tak koi receiver (<-c) ready na ho.
	- Receive (<-c) tab tak BLOCK rahega jab tak koi sender (c <- ...) ready na ho.
So every value transfer is like a “handshake”: sender+receiver same time meet.

----------------------------------------
Step-by-step execution (visual timeline)
----------------------------------------

1) main() starts
	 c := make(chan int)
	 quit := make(chan int)

	 Note: make(chan int) without a size => unbuffered channel.

2) Anonymous goroutine starts:

	 go func() {
			 for i := 0; i < 10; i++ {
					 println(<-c)
			 }
			 quit <- 0
	 }()

	 - Immediately it reaches println(<-c)
	 - That "<-c" is a RECEIVE, so it BLOCKS until fibonacci sends something.
	 - So this goroutine becomes “waiting receiver on c”.

3) main goroutine calls fibonacci(c, quit)

	 fibonacci sets:
		 x = 0, y = 1
	 then loops forever:

		 for {
				 select {
				 case c <- x:
						 x, y = y, x+y
				 case <-quit:
						 println("quit")
						 return
				 }
		 }

----------------------------------------
How select works here
----------------------------------------

select chooses a case that is “ready”:
	- case c <- x is ready IFF some goroutine is currently waiting to receive from c
	- case <-quit is ready IFF some goroutine has sent (or is ready to send) on quit

Since the anonymous goroutine is already waiting on (<-c), the case (c <- x) is ready.

----------------------------------------
The 10-number handshake loop
----------------------------------------

Each iteration looks like this:

	(1) consumer goroutine is waiting on <-c
	(2) fibonacci select picks case c <- x
	(3) value x is transferred (handshake happens)
	(4) consumer prints the received value
	(5) fibonacci updates x,y to next fibonacci pair

Example first few transfers:
	- Start: x=0, y=1
		send 0, print 0, update -> x=1, y=1
	- send 1, print 1, update -> x=1, y=2
	- send 1, print 1, update -> x=2, y=3
	...
So typically you see: 0 1 1 2 3 5 8 13 21 34

Important: fibonacci “free run” nahi karta.
It produces the next number ONLY when someone is ready to receive.
That’s because c is unbuffered.

----------------------------------------
Stopping via quit
----------------------------------------

After printing 10 numbers, consumer goroutine does:
	quit <- 0

That is also an unbuffered SEND, so it blocks until fibonacci receives from quit.

Now, in fibonacci’s next select:
	- case <-quit becomes ready (sender is waiting)
	- case c <- x is NOT ready (no one is receiving from c anymore)

So select chooses <-quit, prints "quit", returns, and program ends.

----------------------------------------
What about default case?
----------------------------------------

If you uncomment default:
	- select will NOT block.
	- When neither send nor quit receive is ready, default runs.
	- That can create a tight loop (CPU high) printing "no one is ready" repeatedly.
So default should be used carefully (often with timeouts/sleep/backoff).
*/

func fibonacci(c chan int, quit chan int) {
	x, y := 0, 1
	for {
		select {
		case c <- x:
			x, y = y, x+y
		case <-quit:
			println("quit")
			return
			// default:
			// 	println("no one is ready") // default case tab chalta hai jb koi bhi case ready na ho
		}
	}
}

func main() {
	c := make(chan int)
	quit := make(chan int)

	// anonymous goroutine
	go func() {
		for i := 0; i < 10; i++ {
			println(<-c)
		}
		quit <- 0
	}()

	fibonacci(c, quit)
}
