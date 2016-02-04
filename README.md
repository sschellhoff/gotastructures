# gotastructures

gotastructures is a lightweight library for datastructures in golang

gotastructures contains
-stack

more datastructures will be added in the future


## Stack example

~~~Go

package main

import (
	"fmt"
	"github.com/sschellhoff/gotastructures"
)

func main() {
	stack := new(gotastructures.Stack)

	fmt.Println(stack.IsEmpty())

	stack.Push(42)
	stack.Push(1337)

	fmt.Println(stack.IsEmpty())

	element := stack.Peek()
	fmt.Println(element.(int))

	fmt.Println(stack.IsEmpty())

	element = stack.Pop()
	element2 := stack.Pop()
	fmt.Println(element.(int))
	fmt.Println(element2.(int))

	fmt.Println(stack.IsEmpty())
}

/* output:

true
false
1337
false
1337
42
true

*/
