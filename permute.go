package iter

type Permutator[T interface{}] struct {
	Index []*Iterator[T]
}

func Permute[T interface{}](Slices ...[]T) *Permutator[T] {
	result := &Permutator[T]{
		Index: make([]*Iterator[T], 0, len(Slices)),
	}
	for _, s := range Slices {
		result.Index = append(result.Index, Iterate[T](s))
	}
	return result
}

func PermuteIterables[T interface{}](Slices ...Iterable[T]) *Permutator[T] {
	result := &Permutator[T]{
		Index: make([]*Iterator[T], 0, len(Slices)),
	}
	for _, s := range Slices {
		result.Index = append(result.Index, s.Iterate())
	}
	return result
}

func (i *Permutator[T]) Size() int {
	result := 1
	for _, j := range i.Index {
		result *= j.Size()
	}
	return result
}

func (i *Permutator[T]) HasNext() bool {
	return len(i.Index) > 0 && i.Index[0].HasNext()
}

func (i *Permutator[T]) Next() {
	j := len(i.Index)-1
	if j >= 0 {
		i.Index[j].Next()
		j -= 1
		for j >= 0 && !i.Index[j+1].HasNext() {
			i.Index[j+1].Reset()
			i.Index[j].Next()
			j -= 1
		}
	}
}

func (i *Permutator[T]) Get() []T {
	result := make([]T, 0, len(i.Index))
	for _, j := range i.Index {
		result = append(result, j.Get())
	}
	return result
}

func (i *Permutator[T]) Ptr() []*T {
	result := make([]*T, 0, len(i.Index))
	for _, j := range i.Index {
		result = append(result, j.Ptr())
	}
	return result
}
