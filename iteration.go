package tiiio

type Array[T, U any] []T

func (arr *Array[T,U]) Reduce(initial U, reducer func(U, T) U) U {
	acc := initial
	for _, item := range *arr {
		acc = reducer(acc, item)
	}
	return acc
}


func (arr *Array[T, U]) Map(iterator func(T) U ) *[]U {
	result := []U{}
	for _, item := range *arr {
		result = append(result, iterator(item))
	}
	return &result
}


