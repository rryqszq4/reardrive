package algorithm

func InsertSort(arr []int) (o int) {

	o = 0;
	var i int = 0;
	var j int = 0;
	var size int = len(arr)
	var tmp int

	for i = 1; i < size; i++ {
		tmp = arr[i];

		for j = i; j>0 && tmp < arr[j-1]; j-- {
			arr[j] = arr[j-1];
			arr[j-1] = tmp;
			o++;
		}
	}
	return o;
}