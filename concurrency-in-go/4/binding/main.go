package main

import (
	"bytes"
	"fmt"
	"sync"
)

func main() {
	//adhoc()
	//lexical()
	nosafe()

}

func adhoc() {
	// dataへのアクセスはloopData内でのみ行う。。。という規約
	data := make([]int, 4)

	loopData := func(handleData chan<- int) {
		defer close(handleData)
		for i := range data {
			handleData <- data[i]
		}
	}

	handleData := make(chan int)

	go loopData(handleData)

	for num := range handleData {
		fmt.Println(num)
	}
}

func lexical() {
	chanOwner := func() <-chan int {
		results := make(chan int, 5) // resultsへの書き込みはこのスコープ内でのみ実施 <- 拘束
		go func() {
			defer close(results)
			for i := 0; i <= 5; i++ {
				results <- i
			}
		}()
		return results
	}

	consumer := func(results <-chan int) { // 読み込みしかできないようにする
		for result := range results {
			fmt.Printf("Received: %d\n", result)
		}
		fmt.Println("Done receiving")
	}

	results := chanOwner() // resultsには読み込み専用
	consumer(results)
}

func nosafe() {
	printData := func(wg *sync.WaitGroup, data []byte) {
		defer wg.Done()

		var buff bytes.Buffer
		for _, b := range data {
			fmt.Fprintf(&buff, "%c", b)
		}
		fmt.Println(buff.String())
	}

	var wg sync.WaitGroup
	wg.Add(2)

	// dataへのアクセスをLockとって同期することもできるが、
	// 独立な部分に分けて、その独立性をレキシカルスコープで管理するとよい。
	data := []byte("golang")
	// data -> data[:3] + data[3:]
	go printData(&wg, data[:3])
	go printData(&wg, data[3:])

	wg.Wait()
}
