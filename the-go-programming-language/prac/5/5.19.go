package main

import "fmt"

func main() {

	i := magicFn()
	fmt.Println(i)

}

func magicFn() (i int) {
	defer func() {
		recover()
		i = 1
	}()
	panic(1)
}
