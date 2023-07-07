package iter

type Zipper[T interface{}] struct {
	Slices [][]T
	Index int
	Count int
}

func Zip[T interface{}](Slices ...[]T) *Zipper[T] {
	result := &Zipper[T]{
		Slices: Slices, 
		Count: -1,
	}
	for _, s := range Slices {
		if result.Count < len(s) {
			result.Count = len(s)
		}
	}
	return result
}

func ZipIterables[T interface{}](Slice ...Iterable[T]) *Zipper[T] {
	result := &Zipper[T]{
		Slices: make([][]T, 0, len(Slice)),
		Count: -1,
	}
	for _, s := range Slice {
		i := s.Iterate()
		result.Slices = append(result.Slices, i.Slice)
		if result.Count < i.Size() {
			result.Count = i.Size()
		}
	}
	return result
}

func (i *Zipper[T]) Size() int {
	return i.Count
}

func (i *Zipper[T]) HasNext() bool {
	return i.Index < i.Size()
}

func (i *Zipper[T]) Next() {
	i.Index++
}

func (i *Zipper[T]) Get() []T {
	result := make([]T, 0, len(i.Slices))
	for _, slice := range i.Slices {
		if i.Index < len(slice) {
			result = append(result, slice[i.Index])
		}
	}
	return result
}

func (i *Zipper[T]) Ptr() []*T {
	result := make([]*T, 0, len(i.Slices))
	for _, slice := range i.Slices {
		if i.Index < len(slice) {
			result = append(result, &slice[i.Index])
		}
	}
	return result
}
