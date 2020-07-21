package main

import (
	"../src/algorithm"
	"fmt"
)

func ExampleInsertSort()  {
	arr := []int{9, 4, 6, 7, 5, 8, 3, 2, 1, 0}

	fmt.Println(arr)

	o := algorithm.InsertSort(arr)

	fmt.Println(arr)

	fmt.Printf("O: %d, len: %d", o, len(arr))
	// output:
	// [9 4 6 7 5 8 3 2 1 0]
	// [0 1 2 3 4 5 6 7 8 9]
	// O: 37, len: 10
}

func ExampleQuickSort() {

	arr  := []int{24,52,11,94,28,36,14,80}

	/*var l int = 50000000;
	var l2 int = l * 2;
	var i int ;
	for i = 0; i < l; i++ {
		arr = append(arr, rand.Intn(l2))
	}*/

	fmt.Println("original array: ",arr)

	//start := time.Now().UnixNano()

	algorithm.QuickSort(arr)

	//fmt.Println(time.Now().UnixNano() - start)

	fmt.Println("quick sort array: ",arr)

	// output:
	// original array:  [24 52 11 94 28 36 14 80]
	// quick sort array:  [11 14 24 28 36 52 80 94]

}

func ExampleMergeSort() {
	arr := []int{80, 70, 11, 72, 25, 36, 44, 10}

	var arr2[] int

	arr2 = algorithm.MergeSort(arr)

	fmt.Println(arr2)

	// output:
	// [10 11 25 36 44 70 72 80]
}