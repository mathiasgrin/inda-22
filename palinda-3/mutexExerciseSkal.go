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
* For each Goroutine we want to decrement the balance
* by 1.
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
