package tiiio

import (
	"strconv"
	"strings"
)

// push: Adds one or more elements to the end of an array and returns the new length.
func (arr *Array) Push(elements ...any) int {
	*arr = append(*arr, elements...)
	return len(*arr)
}

// pop: Removes the last element from an array and returns that element.
func (arr *Array) Pop() any {
	length := len(*arr)
	if length == 0 {
		return nil
	}
	last := (*arr)[length-1]
	*arr = (*arr)[:length-1]
	return last
}

// shift: Removes the first element from an array and returns that removed element.
func (arr *Array) Shift() any {
	if len(*arr) == 0 {
		return nil
	}
	first := (*arr)[0]
	*arr = (*arr)[1:]
	return first
}

// unshift: Adds one or more elements to the front of an array and returns the new length.
func (arr *Array) Unshift(elements ...any) int {
	*arr = append(elements, *arr...)
	return len(*arr)
}

// splice: Changes the contents of an array by removing or replacing existing elements and/or adding new elements.
func (arr *Array) Splice(start, deleteCount int, items ...any) {
	end := start + deleteCount
	*arr = append((*arr)[:start], append(items, (*arr)[end:]...)...)
}

// sort: Sorts the elements of an array in place and returns the sorted array.
func (arr *Array) Sort(compare func(a, b any) bool) {
	length := len(*arr)
	for i := 0; i < length; i++ {
		for j := 0; j < length-1-i; j++ {
			if !compare((*arr)[j], (*arr)[j+1]) {
				(*arr)[j], (*arr)[j+1] = (*arr)[j+1], (*arr)[j]
			}
		}
	}
}

 
// reverse: Reverses the elements of an array in place.
func (arr *Array) Reverse() {
	length := len(*arr)
	for i := 0; i < length/2; i++ {
		(*arr)[i], (*arr)[length-1-i] = (*arr)[length-1-i], (*arr)[i]
	}
}

// fill: Fills all the elements of an array from a start index to an end index with a static value.
func (arr *Array) Fill(value any, start, end int) {
	for i := start; i < end; i++ {
		if i >= 0 && i < len(*arr) {
			(*arr)[i] = value
		}
	}
}


// concat: Returns a new array that is this array joined with other array(s) and/or value(s).
func (arr *Array) Concat(others ...*Array) *Array {
	result := Array{}
	result = append(result, *arr...)
	for _, other := range others {
		result = append(result, *other...)
	}
	return &result
}

// slice: Returns a shallow copy of a portion of an array.
func (arr *Array) Slice(start, end int) *Array {
	length := len(*arr)
	if start < 0 {
		start += length
	}
	if end < 0 {
		end += length
	}
	return &Array{(*arr)[start:end]}
}

// join: Joins all elements of the array into a string.
func (arr *Array) Join(separator string) string {
	strArr := make([]string, len(*arr))
	for i, item := range *arr {
		strArr[i] = anyToString(item)
	}
	return strings.Join(strArr, separator)
}


// Helper function to convert any type to string
func anyToString(a any) string {
	switch v := a.(type) {
	case string:
		return v
	case int:
		return strconv.Itoa(v)
	case float64:
		return strconv.FormatFloat(v, 'f', 6, 64)
	default:
		return ""
	}
}



// toString: Converts the array to a string using commas as separators.
func (arr *Array) ToString() string {
	return arr.Join(",")
}


