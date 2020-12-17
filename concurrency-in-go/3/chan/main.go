package main

import (
	"bytes"
	"fmt"
	"os"
	"sync"
	"time"
)

func main() {
	//intStreamSample(5)
	//strStreamSample()
	//closeAsBroadcast()
	//bufferChannel()
	//nilChannel()
	responsibility()
}

func intStreamSample(cnt int) {
	intStream := make(chan int)
	go func() {
		defer close(intStream)
		for i := 1; i <= cnt; i++ {
			intStream <- i
		}
	}()

	for integer := range intStream {
		fmt.Printf("%d ", integer)
	}

	fmt.Println()
}

func strStreamSample() {
	strStream := make(chan string)
	go func() {
		strStream <- "Hello from chan"
		strStream <- "and goodbye!"
		time.Sleep(1 * time.Second)
		strStream <- "... still listening? Hi!"
		close(strStream)
	}()

	for s, ok := <-strStream; ok; s, ok = <-strStream {
		fmt.Printf("(%v) %v\n", ok, s)
	}
}

func closeAsBroadcast() {
	begin := make(chan interface{})
	var wg sync.WaitGroup
	for i := 0; i < 4; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			<-begin
			fmt.Printf("%v has begun\n", i)
		}(i)
	}

	fmt.Println("Unblocking goroutines...")
	close(begin)
	wg.Wait()
}

func bufferChannel() {
	var bc chan interface{}
	bc = make(chan interface{}, 4)

	bc <- struct{}{}
	bc <- struct{}{}
	bc <- struct{}{}
	bc <- struct{}{}
	// bc <- struct{}{} // <- deadlock!

	var stdoutBuff bytes.Buffer
	defer stdoutBuff.WriteTo(os.Stdout)

	intStream := make(chan int, 2)
	go func() {
		defer close(intStream)
		defer fmt.Fprintln(&stdoutBuff, "Producer Done.")
		for i := 0; i < 10; i++ {
			fmt.Fprintf(&stdoutBuff, "Sending: %d\n", i)
			intStream <- i
		}
	}()

	for integer := range intStream {
		fmt.Fprintf(&stdoutBuff, "Received %v \n", integer)
	}
}

func nilChannel() {
	//var dataStream chan interface{}
	//dataStream <- struct{}{}
	//<-dataStream
	//close(dataStream)
}

func responsibility() {
	chanOwner := func() <-chan int {
		resultStream := make(chan int, 5) // owner shold init
		go func() {
			defer close(resultStream) // owner shold close
			for i := 0; i <= 5; i++ {
				resultStream <- i // owner writes to chan
			}
		}()
		return resultStream // owner capsulate chan as fn
	}

	resultStream := chanOwner()
	for result := range resultStream {
		fmt.Printf("Received: %d\n", result)
	}
	fmt.Println("Done receiving.")
}
