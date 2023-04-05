package main

import (
	"fmt"
	"sync"
)

type bankAccount struct {
	bank    string
	balance int
}


func incrementBalance(s *bankAccount, wg *sync.WaitGroup) {
	(*s).balance = (*s).balance - 1
	// TODO (Do something important on this line)
}

/*
* For each Goroutine we want to increment the car's
* mileage by 1.
* In total, we want the mileage to be set to 2000,
* however, something is wrong...
 */
func main() {
	numOfGoroutines := 1000
	myAccount := bankAccount{"Handelsbanken", 1000}
	var w sync.WaitGroup
	for i := 0; i < numOfGoroutines; i++ {
		// TODO (Do something important on this line)
		go incrementBalance(// TODO)
	}
	w.Wait()
	fmt.Println(myAccount.balance)
}
