package tiiio

type Array []any

func (arr *Array) Reduce(reducer func(any, any) any, initial any) any {
	acc := initial
	for _, item := range *arr {
		acc = reducer(acc, item)
	}
	return acc
}

func (arr *Array) Map(iterator func(any) any) *Array {
	result := Array{}
	for _, item := range *arr {
		result = append(result, iterator(item))
	}
	return &result
}

func (arr *Array) Filter(iterator func(any) bool) *Array {
	result := Array{}
	for _, item := range *arr {
		if iterator(item) {
			result = append(result, item)
		}
	}
	return &result
}

func (arr *Array) ForEach(iterator func(any)) {
	for _, item := range *arr {
		iterator(item)
	}
}

func (arr *Array) Find(iterator func(any) bool) any {
	for _, item := range *arr {
		if iterator(item) {
			return item
		}
	}
	return nil
}

func (arr *Array) FindIndex(iterator func(any) bool) int {
	for i, item := range *arr {
		if iterator(item) {
			return i
		}
	}
	return -1
}

func (arr *Array) Some(iterator func(any) bool) bool {
	for _, item := range *arr {
		if iterator(item) {
			return true
		}
	}
	return false
}

func (arr *Array) Every(iterator func(any) bool) bool {
	for _, item := range *arr {
		if !iterator(item) {
			return false
		}
	}
	return true
}

func (arr *Array) IndexOf(element any, fromIndex int) int {
	for i := fromIndex; i < len(*arr); i++ {
		if (*arr)[i] == element {
			return i
		}
	}
	return -1
}

func (arr *Array) LastIndexOf(element any, fromIndex int) int {
	for i := fromIndex; i >= 0; i-- {
		if (*arr)[i] == element {
			return i
		}
	}
	return -1
}

func (arr *Array) Includes(element any, fromIndex int) bool {
	return arr.IndexOf(element, fromIndex) != -1
}

func (arr *Array) FlatMap(iterator func(any) *Array) *Array {
	result := Array{}
	for _, item := range *arr {
		flattened := iterator(item)
		result = append(result, (*flattened)...)
	}
	return &result
}

func (arr *Array) Flat(depth int) *Array {
	result := Array{}
	for _, item := range *arr {
		if nestedArr, ok := item.(*Array); ok && depth > 0 {
			flattened := nestedArr.Flat(depth - 1)
			result = append(result, (*flattened)...)
		} else {
			result = append(result, item)
		}
	}
	return &result
}

// Note: For ReduceRight, you may want to reverse the array first or iterate backwards.
func (arr *Array) ReduceRight(reducer func(any, any) any, initial any) any {
	acc := initial
	for i := len(*arr) - 1; i >= 0; i-- {
		acc = reducer(acc, (*arr)[i])
	}
	return acc
}
