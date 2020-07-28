package algorithm

import "fmt"

/**
算法题
给你 n 个非负整数 a1，a2，...，an，每个数代表坐标中的一个点 (i, ai) 。在坐标内画 n 条垂直线，垂直线 i 的两个端点分别为 (i, ai) 和 (i, 0)。找出其中的两条线，使得它们与 x 轴共同构成的容器可以容纳最多的水。

说明：你不能倾斜容器，且 n 的值至少为 2


图中垂线代表数组【1,8,6,2,5,4,8,3,7】.在此情况下，容器能够容纳水（表示蓝色部分）的最大值49
 */

func sample(n []int) (r int) {
	if len(n) < 2 {
		return 0
	}
	start := 0
	r = 0
	height := 0
	tmp :=0

	for i := 0; i < len(n); i++ {
		if i ==0 || n[i] > n[start] {
			start = i
			height = n[start]
		}

		if height > 0 && i > start  {

			if height >= n[i] {
				tmp = n[i]*(i-start)
			}else {
				tmp = height*(i-start)
			}

			if tmp > r {
				r = tmp
			}
		}
	}

	return r
}

func ExampleSample() {
	a := []int{1,8,6,2,5,4,8,3,7}
	r := sample(a)

	fmt.Println(r)
	// output:
	// 49
}
