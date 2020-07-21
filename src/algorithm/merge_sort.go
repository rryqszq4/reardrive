package algorithm

func merge(arr_left[] int, arr_right[] int) (result[] int){
	var i int = 0
	var j int = 0

	for i < len(arr_left) && j < len(arr_right) {

		if arr_left[i] < arr_right[j] {
			result = append(result, arr_left[i])
			i++
		} else {
			result = append(result, arr_right[j])
			j++
		}
	}

	if i < len(arr_left) {
		result = append(result, arr_left[i:]...)
	}

	if j < len(arr_right) {
		result = append(result, arr_right[j:]...)
	}

	return result
}

func MergeSort(arr[] int) [] int {

	var i,k,j int
	var size int
	var arr_left[] int
	var arr_right[] int

	i = 0
	size = len(arr)
	k = size -1

	if size <= 1 {
		return arr
	}

	j = (i + k)/2+1

	arr_left = MergeSort(arr[i:j])

	arr_right = MergeSort(arr[j:])

	return merge(arr_left, arr_right)

}