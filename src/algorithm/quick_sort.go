package algorithm

func partation(arr []int) (r int) {
	var i int = 0;
	var k int = 0;
	var m int = 0;

	var size = len(arr);
	var arr2 []int

	k = size - 1

	if (i+k) % 2 == 1 {
		m = (i+k+1)/2
	}else {
		m = (i+k)/2
	}

	arr2 = []int{arr[i],arr[m],arr[k]}

	InsertSort(arr2)

	var tmp int

	tmp = arr[i]
	for i < k {
		if arr[i] == arr2[1] {
			r = i
		}
		if arr[k] == arr2[1] {
			r = k
		}

		if arr[i] >= arr2[1] && arr[k] <= arr2[1]{
			arr[i] = arr[k]
			arr[k] = tmp
			i++
			k--
			tmp = arr[i]
			continue
		}

		if arr[k] > arr2[1] {
			k--
			continue
		}

		if arr[i] < arr2[1] {
			i++
			tmp = arr[i]
			continue
		}
	}


	return r
}

func QuickSort(arr []int) (o int) {
	o = 0;

	var i int = 0;
	var k int = len(arr)-1;

	if i < k {

		r := partation(arr)

		QuickSort(arr[i:r])

		if r < k {
			QuickSort(arr[r:k+1])
		}

	}

	return o;

}

func QuickSort2(arr []int) {
	if len(arr) < 2 {
		return
	}

	i, j := 0, len(arr)-1
	target := arr[0]

	for i < j {
		for arr[j] >= target && j > i {
			j--
		}
		arr[i], arr[j] = arr[j], arr[i]

		for arr[i] < target && i < j {
			i++
		}
		arr[i], arr[j] = arr[j], arr[i]
	}

	QuickSort2(arr[:i])
	QuickSort2(arr[i+1:])
}