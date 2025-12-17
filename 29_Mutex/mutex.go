package main

import (
	"fmt"
	"sync"
)

type post struct {
	views int
	my    sync.Mutex
}

func (p *post) inc(wg *sync.WaitGroup) {
	defer func() {
		p.my.Unlock()
		wg.Done()
	}()

	// modification ke pehle lock and uske baad unlock krna pdega

	p.my.Lock() //lock usi line ko karo jaha modificatoin ho rha hai to avoid bottle neck
	p.views += 1
	// p.my.Unlock() //best practise is to put this into defer
}

// jb bhi concurrent operations hote hai to multiple goroutines work kr rhe hote hai to koi bhi goroutine hoga
// wo value read karega then add krega 1 and at the end write kr dega.

// yahi hota hai yaha pr agar value write honi hoti hai to pehle jisne bhi write kiya hoga wo rewrite krdeta hai
//isiliye mutex help krte hai ki jo bhi common shared variable hai use thoda der tk doosre ko read krne se allow krne
// se thoda time tk lock krke rakhte hai

func main() {
	var wg sync.WaitGroup
	myPost := post{views: 0}

	for i := 0; i < 100; i++ {

		wg.Add(1) // ye sb abhi tk synchronize krne ke liye kiya gya tha
		go myPost.inc(&wg)

	}

	wg.Wait()
	fmt.Println(myPost.views)

}
