package iter

type Iterable[T interface{}] interface {
	Iterate() *Iterator[T]
	IterSize() int
}

func MaxSize[T interface{}](i ...Iterable[T]) int {
	result := 0
	for _, j := range i {
		test := j.IterSize()
		if test > result {
			result = test
		}
	}
	return result
}

func MinSize[T interface{}](i ...Iterable[T]) int {
	result := -1
	for _, j := range i {
		test := j.IterSize()
		if result < 0 || test < result {
			result = test
		}
	}
	return result
}

func SumSize[T interface{}](i ...Iterable[T]) int {
	result := 0
	for _, j := range i {
		result += j.IterSize()
	}
	return result
}

func MulSize[T interface{}](i ...Iterable[T]) int {
	result := 1
	for _, j := range i {
		result *= j.IterSize()
	}
	return result
}

type Iterator[T interface{}] struct {
	Slice []T
	Index int
}

func Iterate[T interface{}](Slice []T) *Iterator[T] {
	return &Iterator[T]{
		Slice: Slice,
		Index: 0,
	}
}

func (i *Iterator[T]) Size() int {
	return len(i.Slice)
}

func (i *Iterator[T]) HasNext() bool {
	return i.Index < len(i.Slice)
}

func (i *Iterator[T]) Next() {
	i.Index++
}

func (i *Iterator[T]) Reset() {
	i.Index = 0
}

func (i *Iterator[T]) Get() T {
	return i.Slice[i.Index]
}

func (i *Iterator[T]) Ptr() *T {
	return &i.Slice[i.Index]
}
