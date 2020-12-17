package main

import (
	"fmt"
	"net/http"
)

func main() {
	//errSample()
	solution()

}

func errSample() {
	checkStatus := func(
		done <-chan interface{},
		urls ...string,
	) <-chan *http.Response {
		responses := make(chan *http.Response)
		go func() {
			defer close(responses)
			for _, url := range urls {
				resp, err := http.Get(url)
				if err != nil {
					fmt.Println(err)
					continue
				}
				select {
				case <-done:
					return
				case responses <- resp:
				}
			}
		}()
		return responses
	}

	done := make(chan interface{})
	defer close(done)

	urls := []string{"https://www.google.com", "https://badhost"}
	for response := range checkStatus(done, urls...) {
		fmt.Printf("Resp: %v\n", response.Status)
	}
}

func solution() {

	type Result struct {
		Error    error
		Response *http.Response
	}
	checkStatus := func(done <-chan interface{}, urls ...string) <-chan Result {
		results := make(chan Result)
		go func() {
			defer close(results)

			for _, url := range urls {
				var result Result
				resp, err := http.Get(url)
				result = Result{Error: err, Response: resp}
				select {
				case <-done:
					return
				case results <- result:
				}
			}

		}()
		return results
	}

	done := make(chan interface{})
	defer close(done)

	urls := []string{"https://www.google.com", "https://badhost"}
	for result := range checkStatus(done, urls...) {
		if result.Error != nil {
			fmt.Printf("Err: %v\n", result.Error)
			continue
		}
		fmt.Printf("Resp: %v\n", result.Response.Status)
	}

	done2 := make(chan interface{})
	defer close(done2)

	urls2 := []string{"https://www.google.com", "https://badhost", "https://a", "https://b", "https://c", "https://d"}
	var errCnt int
	for res := range checkStatus(done2, urls2...) {
		if res.Error != nil {
			fmt.Printf("Err: %v\n", res.Error)
		} else {
			fmt.Printf("Resp: %v\n", res.Response.Status)
			continue
		}
		errCnt++
		if errCnt >= 3 {
			fmt.Println("TOO MANY ERRS: ABORT")
			break
		}
	}

}
