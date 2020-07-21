package algorithm

func SplitAlpha1(x []int, dic map[int][]int) (r [][]int) {

	rr := [][]int{}

	for k,v := range dic {

		if len(rr) <= 0 {
			rr = [][]int{{k}}
			rr[0] = append(rr[0], v...)
			continue
		}

		for i := 0; i < len(rr); i++ {
			tmp := intersect(rr[i], append(v, k))
			if len(tmp) > 0 {
				rr[i] = union(rr[i], append(v, k))
				continue
			}
		}

		rr = append(rr, append(v,k))
	}

	for i := 0; i < len(rr); i++ {
		_is := false
		if len(r) <= 0 {
			r = append(r, rr[i])
			continue
		}
		for j := 0; j < len(r); j++ {
			tmp := intersect(r[j], rr[i])
			if len(tmp) > 0 {
				r[j] = union(r[j], rr[i])
				_is = true
			}
		}
		if !_is {
			r = append(r, rr[i])
		}

	}

	return r
}

func inArray(items []int, item int) bool {
	for _, eachItem := range items {
		if eachItem == item {
			return true
		}
	}
	return false
}

//求并集
func union(slice1, slice2 []int) []int {
	m := make(map[int]int)
	for _, v := range slice1 {
		m[v]++
	}

	for _, v := range slice2 {
		times, _ := m[v]
		if times == 0 {
			slice1 = append(slice1, v)
		}
	}
	return slice1

}

// 求交集
func intersect(slice1 []int, slice2 []int) []int {
	m := make(map[int]int)
	nn := make([]int, 0)
	for _, v := range slice1 {
		m[v]++
	}

	for _, v := range slice2 {
		times, _ := m[v]
		if times == 1 {
			nn = append(nn, v)
		}
	}
	return nn
}

//求差集 slice1-并集
func difference(slice1, slice2 []int) []int {
	m := make(map[int]int)
	nn := make([]int, 0)
	inter := intersect(slice1, slice2)
	for _, v := range inter {
		m[v]++
	}

	for _, value := range slice1 {
		times, _ := m[value]
		if times == 0 {
			nn = append(nn, value)
		}
	}
	return nn
}
