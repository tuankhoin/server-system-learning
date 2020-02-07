package main

import "fmt"

// fibonacci is a function that returns
// a function that returns an int.
func fibonacci() func() int {
	pre := 0
	next := 1
	return func() int {
		sum:=pre+next
		pre=next
		next=sum
		return sum
	}
}

func main() {
	f := fibonacci()
	fmt.Println("0\n1") // I'm lazy, lol
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}
}
