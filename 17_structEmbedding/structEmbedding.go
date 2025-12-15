package main

import (
	"fmt"
	"time"
)

type customer struct {
	name    string
	address string
	mobile  int
}

type order struct {
	id        string
	amount    float32
	status    string
	createdAt time.Time
	customer  // now we are using the abouve struct in this order bucket
}

func main() {
	newOrder := order{
		id:     "3",
		amount: 34,
		status: "pending",
	}

	fmt.Println(newOrder.customer) // when we haven't given anything in this
	fmt.Println(newOrder)

	//lets create new customer
	newCustomer := customer{
		name:    "Shaurya",
		address: "Station Road Khurhand , Banda",
		mobile:  9724562704,
	}

	fmt.Println(newCustomer)

	//using this customer in the order struct => composition 
	secondOrder := order{
		id:     "3",
		amount: 34,
		status: "pending",
		// customer: newCustomer, // used as seperate customer
		customer: customer{ // inline sharing if there is just use of one object of customer
			"Sanskriti Jaiswal",
			"Varanasi",
			9457273940,
		},
	}
	fmt.Println(secondOrder)

}
