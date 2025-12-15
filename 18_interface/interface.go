package main

import "fmt"

type paymenter interface { //convention ye hai ki jb bhi interface banate hai tb tb hume 'er' word last me lagana pdta hai 
	pay(amount float32)
	//if we want to add other methods then to make sure that implementation is working 
	// we also have to add these methods in that struct 
	refund(withdraw float32, amount float32)
}

type payment struct {
	// yaha payment ki configurations wagairah rahengi
	gateway paymenter
	// gateway stripe
}
type razorpay struct {
	// accountNo string
}
type stripe struct {
	// accountNo string
}

// it violates the OPEN CLOSE principle => we are open to extension but not for modification in the code
func (p payment) makePayment(amount float32) {
	// razorpayPaymentGw := razorpay{}
	// razorpayPaymentGw.pay(amount)
	// agar yaha pr hume stripe ka payment gateway use krna hai to yaha pr code ko change krna pdega
	// but instead of changing code => we have to settle this up with the help of interface

	// stripePaymentGw := stripe{}
	p.gateway.pay(amount)
}

func (r razorpay) pay(amount float32) {
	//logic to make payments
	fmt.Println("Making Payment using Razorpay", amount)
}

// lets say we are going to make payment using stripe
func (s stripe) pay(amount float32) {
	fmt.Println("Making Payment using Stripe", amount)
}

// let's say if i want to implement paypal payment gateway
type paypal struct{}

func (p paypal) pay(amount float32){
	fmt.Println("PayPal is processing your payment", amount)
}
func (p paypal) refund(withdraw float32, amount float32){
	fmt.Println("She has withdrawn :",withdraw, " and total amount is : " , amount)
}


func main() {
	// razorpayPGw := razorpay{}
	// stripePGw := stripe{}
	payPalPGw := paypal{}
	paymentInstance := payment{
		gateway: payPalPGw, // this will show error because we want to use payPalGw without adding this refund() function 
	}
	paymentInstance.makePayment(7150)
	
}
