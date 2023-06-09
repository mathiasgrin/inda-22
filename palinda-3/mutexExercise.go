package main

import (
	"fmt"
	"sync"
)

type bankAccount struct {
	bank    string
	balance int
}

func incrementBalance(s *bankAccount, wg *sync.WaitGroup, mtx *sync.Mutex) {
	mtx.Lock()
	(*s).balance = (*s).balance - 1
	mtx.Unlock()
	wg.Done()
}

/*
* For each Goroutine we want to decrement the balance
* by 1.
 */
func main() {
	numOfGoroutines := 1000
	myAccount := bankAccount{"Handelsbanken", 1000}
	var w sync.WaitGroup
	var m sync.Mutex
	for i := 0; i < numOfGoroutines; i++ {
		w.Add(1)
		go incrementBalance(&myAccount, &w, &m)
	}
	w.Wait()
	fmt.Println(myAccount.balance)
}
