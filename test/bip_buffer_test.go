package main

import (
	"../src/core"
	"fmt"
)

func ExampleBipBuffer()  {

	cb := core.NewBipBuffer(32)
	fmt.Println(cb.Size())

	fmt.Println(cb.IsEmpty())

	cb.Offer([]byte("abcdefghijklmnopqrstuvwxyz"))

	//fmt.Println(string(cb.Poll(20)))

	cb.Print()

	cb.Offer([]byte("123456"))

	cb.Print()

	fmt.Println(string(cb.Poll(20)))

	fmt.Println(string(cb.Peek(10)))

	cb.Offer([]byte("1234567890"))

	cb.Print()

	fmt.Println(string(cb.Poll(12)))

	cb.Print()

	fmt.Println(string(cb.Poll(12)))

	cb.Print()

	// output:
	// Size : 32
	// IsEmpyt : true
}