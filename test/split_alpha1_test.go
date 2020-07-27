package main

import (
	"reardrive/src/algorithm"
	"fmt")

func ExampleSplitAlpha1() {
	var x = []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	var dic = map[int][]int{
		1:  {2},
		2:  {1, 4},
		3:  {5},
		4:  {2, 10},
		5:  {3},
		6:  {7},
		7:  {6, 8},
		8:  {7, 9},
		9:  {8},
		10: {4},
	}


	arr := algorithm.SplitAlpha1(x, dic)
	fmt.Println(arr)
	// output:
	// [[1 2 4 10] [3 5] [6 7 8 9]]

}