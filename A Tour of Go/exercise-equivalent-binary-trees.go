package main

import (
	"golang.org/x/tour/tree"
	"fmt"
)

// Walk walks the tree t sending all values
// from the tree to the channel ch.

// In-order traversal
func Walk(t *tree.Tree, c chan int){
	if t.Left != nil {Walk(t.Left, c)}
	c <- t.Value
	if t.Right != nil {Walk(t.Right, c)}
}

func WalkClose(t *tree.Tree, c chan int){
	Walk(t, c)
	close(c)
}

// Same determines whether the trees
// t1 and t2 contain the same values.
func Same(t1, t2 *tree.Tree) bool{
	c1, c2 := make(chan int), make(chan int)
	go WalkClose(t1,c1)
	go WalkClose(t2,c2)
	for i:=range c1 {
		if i != <-c2 {
			return false
		}
	}
	return true
}

func main() {
	ch := make(chan int)
	go WalkClose(tree.New(1), ch)
	for i:= range ch {
		fmt.Print(i, " ")
	}
	fmt.Print("\n")
	fmt.Println(Same(tree.New(1), tree.New(1)), Same(tree.New(1), tree.New(2)))
}
