package core

import "reflect"

func ArrayStringRevers(arr []string) []string{
	for left, right :=0, len(arr)-1; left < right; left, right = left+1, right-1 {
		arr[left], arr[right] = arr[right], arr[left]
	}
	return arr
}

func ArrayRevers(arr []interface{}) []interface{} {
	for left, right := 0, len(arr)-1; left < right; left, right = left+1, right-1 {
		arr[left],arr[right] = arr[right], arr[left]
	}
	return arr
}

func Typeof(v interface{}) string {
	return reflect.TypeOf(v).String()
}