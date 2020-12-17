package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var memoryAccess sync.Mutex
	var value int

	go func() {
		memoryAccess.Lock()
		fmt.Printf("Start sleep::1\n")
		time.Sleep(time.Second)
		fmt.Printf("End sleep::1\n")
		value++
		memoryAccess.Unlock()
	}()
	fmt.Printf("Start sleep::2\n")
	time.Sleep(time.Second)
	fmt.Printf("End sleep::2\n")

	memoryAccess.Lock()
	if value == 0 {
		fmt.Printf("the value is %v\n", value)
	} else {
		fmt.Printf("the value is %v\n", value)
	}
	memoryAccess.Unlock()
}
