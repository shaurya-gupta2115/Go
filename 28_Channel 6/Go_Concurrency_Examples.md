# Go Concurrency - Examples (Hinglish)

Ye file Go concurrency ke practical examples cover karti hai (goroutine, channels, buffered/unbuffered, deadlock, WaitGroup, directional channels) — chhote example se thoda bade example tak.

---

## 21. Simple Goroutine - bina WaitGroup

```go
package main

import (
    "fmt"
    "time"
)

func printMsg() {
    fmt.Println("goroutine se hello")
}

func main() {
    go printMsg()           // background me chal raha hai
    time.Sleep(time.Second) // sirf demo ke liye wait, warna program turant exit ho sakta hai
}
```

- main bhi ek goroutine hai.
- Jaise hi main() khatam, poora program exit, chahe baaki goroutines bach bhi rahi ho.

---

## 22. Same Example - WaitGroup ke saath (better)

```go
package main

import (
    "fmt"
    "sync"
)

func printMsg(wg *sync.WaitGroup) {
    defer wg.Done()
    fmt.Println("goroutine se hello")
}

func main() {
    var wg sync.WaitGroup
    wg.Add(1)

    go printMsg(&wg)

    wg.Wait() // yaha tak tab tak nahi niklega jab tak goroutine Done() na kare
}
```

Yaha hum time.Sleep pe depend nahi, proper synchronization kar rahe hain.

---

## 23. Unbuffered Channel - Deadlock Example

```go
package main

import "fmt"

func main() {
    ch := make(chan string) // unbuffered

    ch <- "ping"           // yaha main goroutine block ho jata hai
    msg := <-ch            // yaha kabhi pohonchta hi nahi

    fmt.Println(msg)
}
```

Runtime error:

```text
fatal error: all goroutines are asleep - deadlock!
```

Reason:

- send ke waqt koi receiver ready nahi
- sirf ek goroutine (main) sab kuch karne ki koshish kar raha hai

---

## 24. Fix 1 - Second Goroutine

```go
package main

import "fmt"

func main() {
    ch := make(chan string)

    go func() {
        ch <- "ping" // background goroutine send karega
    }()

    msg := <-ch       // main goroutine receive karega
    fmt.Println(msg)
}
```

Ab 2 goroutines hain, handshake complete, deadlock nahi.

---

## 25. Fix 2 - Buffered Channel

```go
package main

import "fmt"

func main() {
    ch := make(chan string, 1) // buffer size 1

    ch <- "ping"              // shelf me rakh diya, turant block nahi
    msg := <-ch
    fmt.Println(msg)
}
```

Buffered channel = shelf ya basket: send ke time receiver ka ready hona zaroori nahi (jab tak capacity khali hai).

---

## 26. Buffered Channel + Worker Goroutine (Email Example)

```go
package main

import (
    "fmt"
    "sync"
)

func emailSender(emailChan <-chan string, wg *sync.WaitGroup) {
    defer wg.Done()

    for email := range emailChan { // channel close hone par loop khatam
        fmt.Println("Sending email to:", email)
    }
}

func main() {
    emailChan := make(chan string, 100) // buffered queue
    var wg sync.WaitGroup

    wg.Add(1)
    go emailSender(emailChan, &wg)

    for i := 0; i < 5; i++ {
        emailChan <- fmt.Sprintf("user%d@example.com", i)
    }

    close(emailChan) // signal: ab aur emails nahi aayengi
    wg.Wait()
}
```

Pattern:

- Producer = main goroutine (emails channel me daal raha hai)
- Consumer = emailSender (channel se read karke process kar raha hai)
- WaitGroup = ensure karta hai ki worker complete ho, phir program exit

---

## 27. Directional Channels (chan<- / <-chan)

Type level pe hi bata do kaun send karega, kaun receive:

```go
package main

import "fmt"

// sirf send karne wala
func producer(out chan<- int) {
    for i := 0; i < 3; i++ {
        out <- i
    }
    close(out)
}

// sirf receive karne wala
func consumer(in <-chan int) {
    for v := range in {
        fmt.Println("value:", v)
    }
}

func main() {
    ch := make(chan int, 3)

    go producer(ch)
    consumer(ch)
}
```

- chan<- int  = send-only
- <-chan int  = receive-only
- chan int    = full-duplex (send + receive)

Isse compiler bhi help karta hai galti pakadne me (jaise consumer me accidentally send kar diya ho).

---

### End

Agar ye examples clear ho gaye, tum goroutines + channels + WaitGroup ko confidently use kar paoge.
