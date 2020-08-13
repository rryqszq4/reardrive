package main

import (
	"fmt"
	"reardrive/src/containers"
)

func ExampleCListOne() {
	var list *containers.CList
	var data interface{}

	list = containers.NewCList()

	for i := 0; i < 10; i++ {
		//*data = i

		list.NextInsert(nil, i)
	}

	list.Print()

	fmt.Println("Iterating and removing the fourth element")
	item := list.Head()
	item = list.Next(item)
	item = list.Next(item)

	list.NextRemove(item, &data)

	fmt.Printf("item=%d, %d\n", data, 0)

	list.Print()


	// output:
	//
}