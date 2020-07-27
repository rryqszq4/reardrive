package core

import "reflect"

func ArrayStringRevers(arr []string) []string{
	for left, right :=0, len(arr)-1; left < right; left, right = left+1, right-1 {
		arr[left], arr[right] = arr[right], arr[left]
	}
	return arr
}

func Typeof(v interface{}) string {
	return reflect.TypeOf(v).String()
}

func SliceInterface(s interface{}) (r []interface{}) {
	rs := reflect.ValueOf(s)
	kind := rs.Kind()
	if kind != reflect.Slice && kind != reflect.Array {
		panic(&reflect.ValueError{Method: "utils.SliceInterface", Kind: kind})
	}

	for i := 0; i < rs.Len(); i++ {
		r = append(r, rs.Index(i).Interface())
	}

	return
}

func ArrayIndex(niddle, s interface{}) int {
	slice := SliceInterface(s)

	for k, v := range slice {
		if reflect.DeepEqual(niddle, v) {
			return k
		}
	}

	return -1
}
