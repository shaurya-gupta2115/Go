package main

import "fmt"

// enumerated types
// type OrderStatus int
// const (
// 	Received OrderStatus = iota //iota use hoti haii numerical value ke liye 
// 	Confirmed 
// 	Prepared
// 	Delivered
// )


type OrderStatus string

//  Problem 1: Comma , allowed hi nahi hai
//Type missing in later constants
const (
	Received OrderStatus = "received" //iota use hoti haii numerical value ke liye 
	Confirmed = "confirmed"
	Prepared= "prepared"
	Delivered = "delivered"
)

func changeOrderStatus(status OrderStatus){ //✔ Strong typing & Invalid status pass nahi hoga when we use OrderStatus
	fmt.Println("Changing order status to : " , status)
}

func main(){
	changeOrderStatus(Delivered)
}

/*

package main
import "fmt"
type OrderStatus int

const (
    Received OrderStatus = iota
    Confirmed
    Prepared
    Delivered
)

func (s OrderStatus) String() string {
    switch s {
    case Received:
        return "received"
    case Confirmed:
        return "confirmed"
    case Prepared:
        return "prepared"
    case Delivered:
        return "delivered"
    default:
        return "unknown"
    }
}

func isValidOrderStatus(status OrderStatus) bool {
    switch status {
    case Received, Confirmed, Prepared, Delivered:
        return true
    default:
        return false
    }
}

func changeOrderStatus(status OrderStatus) {
    if !isValidOrderStatus(status) {
        fmt.Println("Invalid status")
        return
    }
    fmt.Println("Changing order status to:", status)
}

func main() {
    changeOrderStatus(Confirmed)
    changeOrderStatus(OrderStatus(10)) // invalid
}

*/