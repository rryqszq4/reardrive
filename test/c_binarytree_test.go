package main

import (
	"fmt"
	"reardrive/src/containers"
)

func ExampleCBinaryTreeOne() {
	var binaryTree *containers.CBinaryTree

	binaryTree = containers.NewCBinaryTree()

	binaryTree.Insert(20)

	binaryTree.Insert(10)

	binaryTree.Insert(30)

	binaryTree.Insert(15)

	binaryTree.Insert(25)

	binaryTree.Insert(70)

	binaryTree.Insert(80)

	binaryTree.Insert(23)

	binaryTree.Insert(26)

	binaryTree.Insert(5)

	binaryTree.Insert(16)

	binaryTree.Insert(6)

	binaryTree.Insert(7)

	binaryTree.Insert(8)

	fmt.Printf("Tree size is %d\n", binaryTree.Size())

	binaryTree.PreorderPrint(binaryTree.Root())

	binaryTree.InorderPrint(binaryTree.Root())

	binaryTree.PostorderPrint(binaryTree.Root())

	result := binaryTree.PreorderTraversal(binaryTree.Root())
	fmt.Println(result)

	result = binaryTree.InorderTraversal(binaryTree.Root())
	fmt.Println(result)

	result = binaryTree.PostorderTraversal(binaryTree.Root())
	fmt.Println(result)

	// output:
	//
}
