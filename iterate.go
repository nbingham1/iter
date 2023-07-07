package iter

func ToSlice[T interface{}](ptr *T) []T {
	return *(*[]T)(unsafe.Pointer(&struct{
		addr uintptr
		len int
		cap int
	}{ptr,1,1}))
}

type Iterable[T interface{}] interface {
	Iterate() *Iterator[T]
}

func MaxSize[T interface{}](i ...Iterable[T]) int {
	result := 0
	for _, j := range i {
		test := j.Iterate()
		if test.Size() > result {
			result = test.Size()
		}
	}
	return result
}

func MinSize[T interface{}](i ...Iterable[T]) int {
	result := -1
	for _, j := range i {
		test := j.Iterate()
		if result < 0 || test.Size() < result {
			result = test.Size()
		}
	}
	return result
}

func SumSize[T interface{}](i ...Iterable[T]) int {
	result := 0
	for _, j := range i {
		result += j.Iterate().Size()
	}
	return result
}

func MulSize[T interface{}](i ...Iterable[T]) int {
	result := 1
	for _, j := range i {
		result *= j.Iterate().Size()
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
