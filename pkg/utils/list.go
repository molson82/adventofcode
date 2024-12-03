package utils

import "reflect"

func RemoveIndex(slice interface{}, index int) interface{} {
	v := reflect.ValueOf(slice)
	if v.Kind() != reflect.Slice {
		panic("removeIndex: not a slice")
	}

	if index < 0 || index >= v.Len() {
		panic("removeIndex: index out of range")
	}

	// Create a new slice with the same type and length as the original
	newSlice := reflect.MakeSlice(v.Type(), v.Len(), v.Len())
	reflect.Copy(newSlice, v)

	// Create the result slice by appending the elements before and after the specified index
	result := reflect.AppendSlice(newSlice.Slice(0, index), newSlice.Slice(index+1, newSlice.Len()))
	return result.Interface()
}
