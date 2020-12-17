package main

import (
	"fmt"
	"os"
	"strconv"
	"time"
)

var pc [256]byte

func init() {
	for i := range pc {
		pc[i] = pc[i/2] + byte(i&1)
	}
}

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Usage: $ command tar")
		return
	}
	ipt, _ := strconv.Atoi(os.Args[1])
	stt := time.Now()
	fmt.Printf("popCount: %v, time: ", popCount(uint64(ipt)))
	fmt.Printf("%v\n", time.Since(stt).Seconds())
	stt = time.Now()
	fmt.Printf("popCountLoop: %v, time: ", popCountLoop(uint64(ipt)))
	fmt.Printf("%v\n", time.Since(stt).Seconds())
	stt = time.Now()
	fmt.Printf("popCountLowestBit: %v, time: ", popCountLowestBit(uint64(ipt)))
	fmt.Printf("%v\n", time.Since(stt).Seconds())
	stt = time.Now()
	fmt.Printf("popCountRemoveSmalls: %v, time: ", popCountRemoveSmalls(uint64(ipt)))
	fmt.Printf("%v\n", time.Since(stt).Seconds())
}

func popCount(x uint64) int {
	return int(pc[byte(x>>(0*8))] +
		pc[byte(x>>(1*8))] +
		pc[byte(x>>(2*8))] +
		pc[byte(x>>(3*8))] +
		pc[byte(x>>(4*8))] +
		pc[byte(x>>(5*8))] +
		pc[byte(x>>(6*8))] +
		pc[byte(x>>(7*8))])
}

func popCountLoop(x uint64) int {
	sum := 0
	for i := 0; i < 8; i++ {
		sum += int(pc[byte(x>>(i*8))])
	}
	return sum
}

func popCountLowestBit(x uint64) int {
	sum := 0
	for i := 0; i < 64; i++ {
		sum += int((x >> i) & 1)
	}
	return sum
}

func popCountRemoveSmalls(x uint64) int {
	sum := 0
	for x != 0 {
		x = x & (x - 1)
		sum++
	}
	return sum
}
