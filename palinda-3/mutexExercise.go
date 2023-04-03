package main

import (
	"fmt"
	"sync"
)

type car struct {
	company string
	model   string
	mileage int
}

func incrementMileage(s *car, wg *sync.WaitGroup, mtx *sync.Mutex) {
	//mtx.Lock()
	(*s).mileage = (*s).mileage + 1
	//mtx.Unlock()
	wg.Done()
}

/*
* For each Goroutine we want to increment the car's
* mileage by 1.
* In total, we want the mileage to be set to 2000,
* however, something is wrong...
 */
func main() {
	numOfGoroutines := 1000
	myVolvo := car{"Volvo", "XC90", 1000}
	var w sync.WaitGroup
	var m sync.Mutex
	for i := 0; i < numOfGoroutines; i++ {
		w.Add(1)
		go incrementMileage(&myVolvo, &w, &m)
	}
	w.Wait()
	fmt.Println(myVolvo.mileage)
}
