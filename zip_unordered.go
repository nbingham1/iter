package iter

import "sort"

type UnorderedZipper[T interface{}] struct {
	Slices [][]T
	Index int
	Count int
}

func (i *UnorderedZipper[T]) Len() int {
	return len(i.Slices)
}

func (i *UnorderedZipper[T]) Swap(j, k int) {
	i.Slices[j], i.Slices[k] = i.Slices[k], i.Slices[j]
}

func (i *UnorderedZipper[T]) Less(j, k int) bool {
	return len(i.Slices[j]) > len(i.Slices[k])
}

func ZipUnordered[T interface{}](Slices ...[]T) *UnorderedZipper[T] {
	result := &UnorderedZipper[T]{
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

func ZipUnorderedIterables[T interface{}](Slice ...Iterable[T]) *UnorderedZipper[T] {
	result := &UnorderedZipper[T]{
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
	sort.Sort(result)
	return result
}

func (i *UnorderedZipper[T]) Size() int {
	return i.Count
}

func (i *UnorderedZipper[T]) HasNext() bool {
	return i.Index < i.Size()
}

func (i *UnorderedZipper[T]) Next() {
	i.Index++
}

func (i *UnorderedZipper[T]) Get() []T {
	result := make([]T, 0, len(i.Slices))
	for _, slice := range i.Slices {
		if i.Index >= len(slice) {
			return result
		}
		result = append(result, slice[i.Index])
	}
	return result
}

func (i *UnorderedZipper[T]) Ptr() []*T {
	result := make([]*T, 0, len(i.Slices))
	for _, slice := range i.Slices {
		if i.Index >= len(slice) {
			return result
		}
		result = append(result, &slice[i.Index])
	}
	return result
}
