package main

func VetCensored() {
	x := []int{1, 2, 3}
	y := append(x)
	for i := range y {
		println(y[i])
	}
}