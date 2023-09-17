package tiiio

type array[T, U any] []T

func (arr *array[T,U]) reduce(initial U, reducer func(U, T) U) U {
	acc := initial
	for _, item := range *arr {
		acc = reducer(acc, item)
	}
	return acc
}


func (arr *array[T, U]) mapping(iterator func(T) U ) *[]U {
	result := []U{}
	for _, item := range *arr {
		result = append(result, iterator(item))
	}
	return &result
}


