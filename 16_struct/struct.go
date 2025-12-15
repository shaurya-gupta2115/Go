package main

import (
	"fmt"
	"time"
)

type order struct { // it is just like class
	id        string
	amount    float32
	status    string
	createdAt time.Time // precision in nanosecond
}

// reciever type or reciever pointer
//.  object_name struct_type hai jo bhi 
func (o *order) changeStatus(status string){
	o.status = status // yaha pr to dereferencing ka use hona chahiye tha...kyu nhi hua ??
	//it's because ki bydefault ye object ki trh by reference recieve hoga and object referential hota hai to wo yaha change nhi hua hai 
}

func (o order) getAmount() float32{ // yaha as a reference pass kr diya hai and continuosly but get ke liye use krne ki koi need nhi hai agar real me manipulation nhi chaihiey ho to 
	return o.amount
}

// creating constructor for order struct
func newOrder( id string, amount float32, status string) *order {
	// initial setup goes here
	myOrder := order{
		id: id, 
		amount : amount,
		status : status,
		createdAt: time.Now(),
	}
	return &myOrder // convention hai ki hum jo bhi return krenge wo pointer return krenge
	// uske krn jo bhi return type hoga wo ek pointer container hoga 
}


func main() { 
	// var order order=

	// if i don't set the value at any field, then by default value zero-valued rehti hai 

	myOrder := order{
		id:     "hafogSgaG34lghfo@gewr4f",
		amount: 45,
		status: "recieved",
	}

	myOrder.changeStatus("confirmed")
	fmt.Println(myOrder) // yaha pr update ho gya hai ab => answer abhi nhi ...kyunki ye abhi tk pass by value hai 

	fmt.Println(myOrder.getAmount())

	// myOrder.createdAt  = time.Now()
	// fmt.Println(myOrder.status)

	myOrder2 := order{
		id:        "2",
		amount:    2352,
		status:    "delivered",
		createdAt: time.Now(),
	}

	myOrder.status = "paid"
	fmt.Println(myOrder)
	fmt.Println(myOrder2) // both are different instance and do not affect each other 

	fmt.Println()

	//calling constructor function 
	secondOrder := newOrder("5", 2327, "accepted")
	fmt.Println(secondOrder)
	fmt.Println(secondOrder.status)

	//=========================================part =========2===================================

	//creating inline struct other conventional "type nameOfClass struct{}" like this we have to do
	language := struct {
		name string 
		isGood bool
	}{"golang", true}

	fmt.Println(language)

	

}
