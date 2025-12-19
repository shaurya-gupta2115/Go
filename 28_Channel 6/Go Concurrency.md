# Goroutines, Channels & WaitGroups -- Deep Dive (Hinglish)

> **Goal:** Go concurrency ko zero se leke advanced level tak samajhna
> -- goroutines, channels, deadlock, WaitGroup, aur unke real-life
> analogies ke saath.

------------------------------------------------------------------------

## 1. Goroutine kya hoti hai?

**Goroutine = lightweight thread** jo Go runtime manage karta hai.

### Analogy 1: Restaurant Kitchen 🍳

-   Chef = main goroutine\
-   Helpers = goroutines\
-   Sab kaam parallel chal sakta hai

``` go
go cookFood()
```

Bas `go` keyword lagaya aur kaam background me chala gaya.

------------------------------------------------------------------------

## 2. Goroutine kyu use karte hain?

-   Fast
-   Cheap (kam memory)
-   Concurrent kaam

### Analogy 2: Call Center 📞

-   Ek agent ek call handle kare → slow
-   Multiple agents (goroutines) → fast

------------------------------------------------------------------------

## 3. Problem: Goroutine ka result kaise mile?

Goroutine directly `return` nahi kar sakti.

❌ Wrong thinking:

``` go
res := go sum(2,3)
```

✅ Solution: **Channels**

------------------------------------------------------------------------

## 4. Channel kya hota hai?

**Channel = typed pipe** jisse goroutines baat karti hain.

``` go
ch := make(chan int)
```

### Analogy 3: Water Pipe 🚰

-   Tank = sender
-   Pipe = channel
-   House = receiver

------------------------------------------------------------------------

## 5. Send & Receive

``` go
ch <- 10     // send
x := <-ch   // receive
```

### Analogy 4: Courier 📦

-   Sender parcel bhejta
-   Receiver tabhi milega jab parcel aaye

------------------------------------------------------------------------

## 6. Unbuffered Channel (Default)

-   Capacity = 0
-   Send & Receive saath-saath

### Analogy 5: Handshake 🤝

-   Dono ready → tabhi deal complete

------------------------------------------------------------------------

## 7. Buffered Channel

``` go
ch := make(chan int, 2)
```

### Analogy 6: Basket 🧺

-   Basket me 2 fruits rakh sakte ho
-   Uske baad rukna padega

------------------------------------------------------------------------

## 8. Deadlock kya hota hai?

> Jab sab goroutines wait kar rahi ho aur koi progress na ho.

``` text
fatal error: all goroutines are asleep - deadlock!
```

### Analogy 7: Dono chup 🤐

-   A: "Pehle tum bolo"
-   B: "Nahi, pehle tum"

------------------------------------------------------------------------

## 9. Deadlock Example

``` go
func main() {
    ch := make(chan int)
    <-ch // koi send hi nahi kar raha
}
```

------------------------------------------------------------------------

## 10. Channel as Return Value

``` go
func sum(ch chan int, a, b int) {
    ch <- a + b
}
```

### Analogy 8: Result Slip 🧾

-   Clerk kaam karta
-   Slip counter pe rakh deta

------------------------------------------------------------------------

## 11. Completion Signal Channel

``` go
done := make(chan bool)
go task(done)
<-done
```

### Analogy 9: Race Finish Flag 🏁

-   Jab tak flag na gire, race over nahi

------------------------------------------------------------------------

## 12. `defer` + Channel (Safe Pattern)

``` go
defer func() {
    done <- true
}()
```

### Analogy 10: Insurance ☂️

-   Accident ho ya na ho, protection milta hai

------------------------------------------------------------------------

## 13. WaitGroup kya hota hai?

WaitGroup sirf **wait** karta hai.

``` go
var wg sync.WaitGroup
wg.Add(1)
go work(&wg)
wg.Wait()
```

### Analogy 11: Attendance Register 📋

-   Jab tak sab sign na karein, class khatam nahi

------------------------------------------------------------------------

## 14. Channel vs WaitGroup

  Feature         Channel   WaitGroup
  --------------- --------- -----------
  Data pass       ✅        ❌
  Sirf wait       ⚠️        ✅
  Deadlock risk   Yes       No
  Intent clear    Medium    High

------------------------------------------------------------------------

## 15. Kab kya use karein?

### Use Channel when:

-   Data bhejna ho
-   Result chahiye
-   Pipeline / worker pool

### Use WaitGroup when:

-   Sirf wait karna ho
-   Multiple goroutines ka end

### Analogy 12:

-   Channel = WhatsApp
-   WaitGroup = Roll Call

------------------------------------------------------------------------

## 16. Channel + WaitGroup Together

``` go
jobs := make(chan int)
var wg sync.WaitGroup
```

### Analogy 13: Factory 🏭

-   Belt (channel)
-   Supervisor (WaitGroup)

------------------------------------------------------------------------

## 17. Common Mistakes

-   `time.Sleep` use karna ❌
-   Channel close galat jagah ❌
-   `wg.Add` goroutine ke andar ❌

### Analogy 14: Wrong Wiring ⚡

------------------------------------------------------------------------

## 18. Golden Rule ⭐

> **If you wait → WaitGroup**\
> **If you communicate → Channel**

### Analogy 15: Phone vs Alarm

-   Phone (channel) = baat + info
-   Alarm (WaitGroup) = sirf signal

------------------------------------------------------------------------

## 19. Interview One-Liners

-   Goroutine: lightweight thread
-   Channel: communication pipe
-   WaitGroup: synchronization primitive
-   Deadlock: all goroutines waiting

------------------------------------------------------------------------

## 20. Final Mental Model

    [Goroutine] → Channel → [Goroutine]
            |
         WaitGroup

------------------------------------------------------------------------

### ✅ End Note

Agar ye concepts clear ho gaye, Go backend + systems programming
**smooth ho jaata hai** 🚀
